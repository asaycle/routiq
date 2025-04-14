import { TouringServiceClient } from '../../lib/proto/v1/TouringServiceClientPb';
import {
    CreateTouringRequest,
    ListTouringsRequest,
    Touring,
} from '../../lib/proto/v1/touring_pb';
import { Date as PbDate } from "../../lib/proto/google/type/date_pb"; // google-protobuf の Date 型をインポート
import { Tag } from '../../lib/proto/v1/tag_pb';


// gRPC-Webクライアントを初期化
const client = new TouringServiceClient('http://localhost:8080');

const toPbDate = (jsDate: Date): PbDate => {
    const pbDate = new PbDate();
    pbDate.setYear(jsDate.getFullYear());
    pbDate.setMonth(jsDate.getMonth() + 1); // getMonth() は 0 ベースなので +1
    pbDate.setDay(jsDate.getDate());
    return pbDate;
};

export const createTouring = async (
    routeID: string | null,
    title: string,
    date: Date | null,
    score: number = 0,
    tags: Tag[],
): Promise<Touring | null> => {
    const touring = new Touring();
    if (routeID !== null) {
        touring.setRouteId(routeID);
    }
    touring.setTitle(title);
    if (date !== null) {
        touring.setDate(toPbDate(date));
    }
    touring.setScore(score);
    touring.setTagsList(tags);

    const request = new CreateTouringRequest();
    request.setTouring(touring);

    return new Promise((resolve, reject) => {
        client.createTouring(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response || null);
            }
        });
    });
};

export const listRouteTourings = async (routeID: string): Promise<Touring[]> => {
    const request = new ListTouringsRequest();
    request.setFilter(`route_id == "${routeID}"`)
    return new Promise((resolve, reject) => {
        client.listTourings(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response?.getTouringsList() || []);
            }
        });
    });
};

