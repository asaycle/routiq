import { RidingServiceClient } from '../../lib/proto/v1/RidingServiceClientPb';
import {
    CreateRidingRequest,
    ListRidingsRequest,
    Riding,
} from '../../lib/proto/v1/riding_pb';
import { Date as PbDate } from "../../lib/proto/google/type/date_pb"; // google-protobuf の Date 型をインポート
import { Tag } from '../../lib/proto/v1/tag_pb';


// gRPC-Webクライアントを初期化
const client = new RidingServiceClient('http://localhost:8080');

const toPbDate = (jsDate: Date): PbDate => {
    const pbDate = new PbDate();
    pbDate.setYear(jsDate.getFullYear());
    pbDate.setMonth(jsDate.getMonth() + 1); // getMonth() は 0 ベースなので +1
    pbDate.setDay(jsDate.getDate());
    return pbDate;
};

export const createRiding = async (
    routeID: string | null,
    title: string,
    date: Date | null,
    score: number = 0,
    tags: Tag[],
): Promise<Riding | null> => {
    const riding = new Riding();
    if (routeID !== null) {
        riding.setRouteId(routeID);
    }
    riding.setTitle(title);
    if (date !== null) {
        riding.setDate(toPbDate(date));
    }
    riding.setScore(score);
    riding.setTagsList(tags);

    const request = new CreateRidingRequest();
    request.setRiding(riding);

    return new Promise((resolve, reject) => {
        client.createRiding(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response || null);
            }
        });
    });
};

export const listRouteRidings = async (routeID: string): Promise<Riding[]> => {
    const request = new ListRidingsRequest();
    request.setFilter(`route_id == "${routeID}"`)
    return new Promise((resolve, reject) => {
        client.listRidings(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response?.getRidingsList() || []);
            }
        });
    });
};

