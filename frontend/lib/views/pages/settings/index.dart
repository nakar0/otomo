import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/configs/app_themes.dart';
import 'package:otomo/view_models/user.dart';
import 'package:otomo/views/router.dart';
import 'package:settings_ui/settings_ui.dart';

class SettingsPage extends HookConsumerWidget {
  const SettingsPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final appTheme = Theme.of(context).extension<AppTheme>();
    final dangerTitleTextStyle = TextStyle(
      color: appTheme?.dangerColor,
      fontWeight: FontWeight.bold,
    );

    final user = ref.watch(userProvider);
    final userNotifier = ref.read(userProvider.notifier);

    return Scaffold(
      appBar: AppBar(automaticallyImplyLeading: true),
      body: SettingsList(
        lightTheme: appTheme?.settingsTheme,
        sections: [
          SettingsSection(
            title: const Text('Account'),
            tiles: [
              SettingsTile(
                leading: const Icon(Icons.email),
                title: const Text('メールアドレス', style: TextStyle(fontSize: 12)),
                value: Text(user?.email ?? '',
                    style: const TextStyle(fontSize: 14)),
              ),
            ],
          ),

          // # Profile
          // language

          // # About
          // Privacy policy
          // Terms of service
          // Acknowledgements

          // Sign Out
          SettingsSection(
            tiles: [
              SettingsTile(
                title: Text('ログアウト', style: dangerTitleTextStyle),
                onPressed: (_) => userNotifier.signOut(),
              ),
            ],
          ),

          // version

          // # Danger area
          SettingsSection(
            tiles: [
              SettingsTile(
                  title: Text('アカウント削除', style: dangerTitleTextStyle),
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
