import 'package:flutter_chat_types/flutter_chat_types.dart' as chat;
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:otomo/entities/lat_lng.dart';
import 'package:otomo/entities/message.dart';
import 'package:otomo/view_models/boundary/message.dart';

class Converter {
  Converter._();

  static LatLng latLngToGoogle(AppLatLng ll) =>
      LatLng(ll.latitude, ll.longitude);

  static List<chat.TextMessage> textMessageDataToChatTextMessageList(
    List<TextMessageData> messages,
  ) =>
      messages.map(textMessageDataToChatTextMessage).toList();

  static chat.TextMessage textMessageDataToChatTextMessage(
    TextMessageData message,
  ) =>
      chat.TextMessage(
        id: message.id,
        text: message.text,
        author: roleToChatUser(message.role),
        remoteId: message.remoteId,
        createdAt: message.sentAt.millisecondsSinceEpoch,
        status: messageStatusToChatStatus(message.status)
      );

  static chat.User roleToChatUser(Role role) => chat.User(id: role.toString());

  static chat.Status messageStatusToChatStatus(MessageStatus status) {
    switch (status) {
      case MessageStatus.sent:
        return chat.Status.sent;
      case MessageStatus.sending:
        return chat.Status.sending;
      case MessageStatus.error:
        return chat.Status.error;
    }
  }
}
