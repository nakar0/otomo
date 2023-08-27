import 'package:injectable/injectable.dart';
import 'package:otomo/grpc/generated/message.pb.dart';

import '../grpc/generated/chat_service_service.pbgrpc.dart';

@Injectable()
class ChatController {
  ChatController(this._chatService);

  final ChatServiceClient _chatService;

  Future<List<Message>> listMessages(
    String userId,
    int? pageSize,
    String? pageStartAfterMessageId,
  ) async {
    final req = ChatService_ListMessagesRequest()
      ..userId = userId
      ..pageSize = pageSize ?? 0
      ..pageStartAfterMessageId = pageStartAfterMessageId ?? '';
    final resp = await _chatService.listMessages(req);
    return resp.messages;
  }

  Stream<String> sendMessage(
    String text,
  ) {
    final req = ChatService_SendMessageRequest()..text = text;
    return _chatService.sendMessage(req).map((replyChunk) => replyChunk.text);
  }
}
