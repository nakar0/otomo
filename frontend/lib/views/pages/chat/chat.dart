import 'package:flutter/material.dart';
import 'package:flutter_chat_types/flutter_chat_types.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/viewmodels/chat.dart';
import 'package:otomo/views/cases/chat/chat_modal_ui_leading.dart';
import 'package:otomo/views/cases/chat/chat_ui.dart';
import 'package:otomo/views/cases/chat/chat_ui_app_bar.dart';

class ModalChat extends ConsumerWidget {
  const ModalChat({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final chat = ref.watch(chatProvider);
    return Scaffold(
      appBar: const ChatUIAppBar(
        leading: ChatModalUILeading(),
        online: true,
        title: 'Chat',
      ),
      body: ChatUI(
        messages: chat.value?.messages ?? [],
        onSendPressed: (message) async {
          ref.read(chatProvider.notifier).sendMessage(message.text);
        },
        user: chat.value?.user ?? const User(id: ''),
      ),
    );
  }
}
