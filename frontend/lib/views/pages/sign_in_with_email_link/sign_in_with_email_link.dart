import 'package:flutter/material.dart';
import 'package:otomo/views/bases/buttons/rounded_filled_button.dart';
import 'package:otomo/views/bases/layouts/safe_area_layout.dart';
import 'package:otomo/views/bases/layouts/side_space_layout.dart';
import 'package:otomo/views/bases/screens/scaffold_with_only_back_button.dart';
import 'package:otomo/views/bases/spaces/space.dart';
import 'package:otomo/views/bases/spaces/spaces.dart';
import 'package:otomo/views/bases/text_fields/rounded_text_form_field.dart';
import 'package:otomo/views/bases/text_fields/text_field_label.dart';

class SignInWithEmailLinkPage extends StatelessWidget {
  const SignInWithEmailLinkPage({super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return ScaffoldWithOnlyBackButton(
      body: SafeAreaLayout(
        child: SideSpaceLayout(
          child: Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  'Continue with email',
                  style: theme.textTheme.headlineMedium,
                ),
                Spaces.h40,
                const TextFieldLabel(label: 'Email'),
                const RoundedTextFormField(
                  keyboardType: TextInputType.emailAddress,
                  hintText: 'Email',
                ),
                const Space(height: 80),
                RoundedFilledButton.large(
                  verticalExpanded: true,
                  onPressed: null,
                  child: const Text('Continue'),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}