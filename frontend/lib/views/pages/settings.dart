import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/constants/links.dart';
import 'package:otomo/tools/app_package_info.dart';
import 'package:otomo/view_models/account.dart';
import 'package:otomo/view_models/router.dart';
import 'package:otomo/views/bases/texts/texts.dart';
import 'package:otomo/views/cases/danger/danger_text.dart';
import 'package:otomo/views/cases/settings/app_settings_list.dart';
import 'package:otomo/views/cases/settings/app_settings_section.dart';
import 'package:otomo/views/cases/settings/app_settings_tile.dart';
import 'package:otomo/views/routes.dart';
import 'package:otomo/views/utils/launcher.dart';
import 'package:otomo/views/utils/localizations.dart';

class SettingsPage extends HookConsumerWidget {
  const SettingsPage({super.key});

  Widget _buildSettingsList(BuildContext context, WidgetRef ref) {
    final account = ref.watch(accountVMProvider).account;
    final accountNotifier = ref.read(accountVMProvider.notifier);
    return AppSettingsList(
      shrinkWrap: true,
      sections: [
        AppSettingsSection(
          title: Text(context.l10n.settingsPageAccountSection),
          tiles: [
            AppSettingsTile(
              title: BodySmall(context.l10n.email),
              value: Text(account?.email ?? ''),
            ),
          ],
        ),

        // # Profile
        // language

        // # About
        AppSettingsSection(
          title: Text(context.l10n.settingsPageAboutSection),
          tiles: [
            AppSettingsTile(
              title: Text(context.l10n.privacyPolicy),
              trailing: const Icon(Icons.keyboard_arrow_right_rounded),
              onPressed: (_) => Launcher.urlString(Links.privacyPolicy),
            ),
            AppSettingsTile(
              title: Text(context.l10n.terms),
              trailing: const Icon(Icons.keyboard_arrow_right_rounded),
              onPressed: (_) => Launcher.urlString(Links.terms),
            ),
            // Acknowledgements
          ],
        ),

        AppSettingsSection(
          title: Text(context.l10n.settingsPageHelpSection),
          tiles: [
            AppSettingsTile(
              title: Text(context.l10n.contactUs),
              trailing: const Icon(Icons.keyboard_arrow_right_rounded),
              onPressed: (_) => Launcher.inquiry(
                locale: localeOf(context),
                userId: account?.uid ?? '',
                email: account?.email ?? '',
              ),
            ),
          ],
        ),

        AppSettingsSection(
          tiles: [
            AppSettingsTile(
              title: DangerText(context.l10n.signOut),
              onPressed: (_) => accountNotifier.signOut(),
            ),
          ],
        ),

        AppSettingsSection(
          tiles: [
            AppSettingsTile(
                title: DangerText(context.l10n.accountDeletion),
                trailing: const Icon(Icons.keyboard_arrow_right_rounded),
                onPressed: (_) =>
                    ref.read(routerProvider).push(Routes.accountDeletion)),
          ],
        ),
      ],
    );
  }

  Widget _buildVersion(BuildContext context) {
    return BodySmall(
      context.l10n.settingsPageVersion(
          AppPackageInfo.version, AppPackageInfo.buildNumber),
    );
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Scaffold(
      appBar: AppBar(
        title:
            TitleMedium(context.l10n.settingsPageTitle, style: TextStyles.bold),
      ),
      body: Column(
        mainAxisSize: MainAxisSize.min,
        children: [_buildSettingsList(context, ref), _buildVersion(context)],
      ),
    );
  }
}
