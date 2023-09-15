import { Metadata } from '@grpc/grpc-js';
import {
  ChatServiceClient,
} from '../protos/chat_service_grpc_pb';
import { ChatService_AskToMessageRequest } from '../protos/chat_service_pb';

export class ChatController {
  constructor(
    private _client: ChatServiceClient,
    private _metadata: Metadata,
  ) { }

  async askToMessage(userId: string) {
    console.log('called ChatController.askToMessage');

    const request = new ChatService_AskToMessageRequest();
    request.setUserId(userId);

    this._client.askToMessage(
      request,
      this._metadata,
      (err, value) => {
        if (err) {
          console.warn(`error: ${err}`);
        } else {
          console.log(`response: ${value}`);
        }
      }
    );
  }
}
