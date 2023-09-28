import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/configs/links.dart';
import 'package:otomo/view_models/user.dart';
import 'package:otomo/views/bases/texts/texts.dart';
import 'package:otomo/views/cases/danger/danger_text.dart';
import 'package:otomo/views/cases/settings/app_settings_list.dart';
import 'package:otomo/views/cases/settings/app_settings_section.dart';
import 'package:otomo/views/cases/settings/app_settings_tile.dart';
import 'package:otomo/views/router.dart';
import 'package:url_launcher/url_launcher_string.dart';

class SettingsPage extends HookConsumerWidget {
  const SettingsPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(userProvider);
    final userNotifier = ref.read(userProvider.notifier);

    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: true,
        title: const TitleMedium('設定', style: TextStyles.bold),
      ),
      body: AppSettingsList(
        sections: [
          AppSettingsSection(
            title: const Text('アカウント'),
            tiles: [
              AppSettingsTile(
                leading: const Icon(Icons.email),
                title: const BodySmall('メールアドレス'),
                value: Text(user?.email ?? ''),
              ),
            ],
          ),

          // # Profile
          // language

          // # About
          // Privacy policy
          // Terms of service
          // Acknowledgements

          AppSettingsSection(
            title: const Text('ヘルプ'),
            tiles: [
              AppSettingsTile(
                title: const Text('お問い合わせ'),
                leading: const Icon(Icons.question_mark_rounded),
                trailing: const Icon(Icons.keyboard_arrow_right_rounded),
                onPressed: (_) => launchUrlString(
                    Links.inquiry(user?.id ?? '', user?.email ?? '')),
              ),
            ],
          ),

          AppSettingsSection(
            tiles: [
              AppSettingsTile(
                title: const DangerText('ログアウト'),
                onPressed: (_) => userNotifier.signOut(),
              ),
            ],
          ),

          // version

          // # Danger area
          AppSettingsSection(
            tiles: [
              AppSettingsTile(
                  title: const DangerText('アカウント削除'),
                  trailing: const Icon(Icons.keyboard_arrow_right_rounded),
                  onPressed: (_) =>
                      ref.read(routerProvider).push(Routes.accountDeletion)),
            ],
          ),
        ],
      ),
    );
  }
}
