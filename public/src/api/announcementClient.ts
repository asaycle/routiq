import { AnnouncementServiceClient } from '../../lib/proto/v1/AnnouncementServiceClientPb';
import { ListAnnouncementsRequest, ListAnnouncementsResponse, Announcement } from '../../lib/proto/v1/announcement_pb';

// gRPC-Webクライアントを初期化
const client = new AnnouncementServiceClient(process.env.REACT_APP_API_BASE!);


export const listAnnouncements = async (filter: string, itemSize: number): Promise<Announcement[]> => {
    const request = new ListAnnouncementsRequest();
    request.setFilter(filter);
    request.setPageSize(itemSize);
    return new Promise((resolve, reject) => {
        client.listAnnouncements(request, {}, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response?.getAnnouncementsList() || []);
            }
        });
    });
};

