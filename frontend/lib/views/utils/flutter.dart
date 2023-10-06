import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

final class FlutterUtils {
  FlutterUtils._();
  static void afterBuildCallback(VoidCallback callback) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      callback();
    });
  }

  static void unfocus(BuildContext context) {
    final currentScope = FocusScope.of(context);
    if (!currentScope.hasPrimaryFocus && currentScope.hasFocus) {
      FocusManager.instance.primaryFocus?.unfocus();
    }
  }

  static Future<void> hideKeyboard() =>
      SystemChannels.textInput.invokeMethod('TextInput.hide');

  static Future<void> copyText(String text) =>
      Clipboard.setData(ClipboardData(text: text));

  static bool validateAndSaveForm(GlobalKey<FormState> formKey) {
    assert(formKey.currentState != null);

    if (!(formKey.currentState?.validate() ?? false)) return false;
    formKey.currentState?.save();
    return true;
  }
}
