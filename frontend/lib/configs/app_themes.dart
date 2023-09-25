import 'package:flutter/material.dart';
import 'package:flutter_chat_ui/flutter_chat_ui.dart';
import 'package:otomo/configs/app_colors.dart';
import 'package:otomo/views/bases/text_fields/rounded_text_form_field.dart';
import 'package:settings_ui/settings_ui.dart';

abstract class AppThemes {
  static const _lightSecondary = Color(0xFFC3E874);
  static const _lightBackground = Color(0xFF272729);
  static const _lightBottomNavIconDisabled = Color(0xFFA2A7B1);
  static final ThemeData _light = ThemeData.from(
    colorScheme: const ColorScheme.light(
      brightness: Brightness.light,
      primary: AppColors.primary,
      onPrimary: Colors.white,
      secondary: _lightSecondary,
      onSecondary: Colors.white,
      background: Colors.white,
      onBackground: _lightBackground,
    ),
    useMaterial3: true,
  );

  static final ThemeData light = _light.copyWith(
    inputDecorationTheme: const InputDecorationTheme(
      fillColor: Color(0xFFE9E9EA),
      hintStyle: TextStyle(
        color: Color(0xFFB5B5B8),
      ),
    ),
    appBarTheme: AppBarTheme(
      backgroundColor: _light.colorScheme.background,
      shadowColor: Colors.black12,
    ),
    bottomNavigationBarTheme: BottomNavigationBarThemeData(
      backgroundColor: _light.colorScheme.background,
      selectedItemColor: _light.colorScheme.primary,
      selectedIconTheme: IconThemeData(
        size: 28,
        color: _light.colorScheme.primary,
      ),
      unselectedItemColor: _lightBottomNavIconDisabled,
      unselectedIconTheme: const IconThemeData(
        size: 28,
        color: _lightBottomNavIconDisabled,
      ),
    ),
    extensions: <ThemeExtension>[
      AppTheme(
        dangerColor: const Color(0xFFE74C3C),
        chatTheme: DefaultChatTheme(
          primaryColor: _light.colorScheme.primary,
          secondaryColor: const Color(0xFFF7F7F7),
          backgroundColor: _light.colorScheme.background,
          sendButtonIcon:
              Icon(Icons.send_rounded, color: _light.colorScheme.primary),
          inputTextStyle: _light.textTheme.bodyMedium!,
          inputTextColor: _light.textTheme.bodyMedium!.color!,
          inputBorderRadius: BorderRadius.zero,
          inputBackgroundColor: _light.colorScheme.background,
          inputContainerDecoration: BoxDecoration(
            border: Border(
              top: BorderSide(color: _light.dividerColor),
            ),
          ),
          inputTextDecoration: RoundedTextFormField.inputDecoration,
          messageBorderRadius: 16,
        ),
        settingsTheme: const SettingsThemeData(
          settingsSectionBackground: Color(0xFFF7F7F7),
          settingsListBackground: Colors.white,
        ),
      )
    ],
  );

  // static final ThemeData oldLight = ThemeData(
  //   inputDecorationTheme: InputDecorationTheme(),
  //   textTheme: Typography.material2021(colorScheme: ColorScheme.light()).black,

  //   // Fixed
  //   useMaterial3: true,
  // );
  static final ThemeData dark = ThemeData.from(
    colorScheme: const ColorScheme.dark(),
    useMaterial3: true,
    // bottomNavigationBarTheme: const BottomNavigationBarThemeData(
    //   backgroundColor: AppColors.lightBlack,
    //   selectedItemColor: AppColors.white,
    //   selectedIconTheme: IconThemeData(
    //     size: 28,
    //     color: AppColors.lightBlack,
    //   ),
    //   unselectedItemColor: AppColors.lightBlack,
    //   unselectedIconTheme: IconThemeData(
    //     size: 28,
    //     color: AppColors.white,
    //   ),
    // ),
  );
}

class AppTheme extends ThemeExtension<AppTheme> {
  AppTheme({
    required this.chatTheme,
    required this.settingsTheme,
    required this.dangerColor,
  });

  final ChatTheme chatTheme;
  final SettingsThemeData settingsTheme;

  final Color dangerColor;

  @override
  ThemeExtension<AppTheme> copyWith() {
    return this;
  }

  @override
  ThemeExtension<AppTheme> lerp(
    covariant ThemeExtension<AppTheme>? other,
    double t,
  ) {
    return this;
  }
}
