import { TagServiceClient } from '../../lib/proto/v1/tag_grpc_web_pb';
import { CreateTagRequest, ListTagsRequest, Tag } from '../../lib/proto/v1/tag_pb';

// gRPC-Webクライアントを初期化
const client = new TagServiceClient(process.env.REACT_APP_API_BASE!);

export const createTag = async (
  name: string,
): Promise<Tag> => {
  const tag = new Tag();
  tag.setName(name);
  const request = new CreateTagRequest();
  request.setTag(tag);

  return new Promise((resolve, reject) => {
    client.createTag(request, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response || null);
      }
    });
  });
};

/**
 * ListTags APIを呼び出す関数
 */
export const listTags = async (filter: string): Promise<Tag[]> => {
  const request = new ListTagsRequest();
  request.setFilter(filter);
  request.setPageSize(10);
  return new Promise((resolve, reject) => {
    client.listTags(request, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response?.getTagsList() || []);
      }
    });
  });
};
