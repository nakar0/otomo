import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/view_models/splash.dart';
import 'package:otomo/views/bases/screens/scaffold_with_barrier_indicator.dart';
import 'package:otomo/views/cases/logo/text_logo.dart';

class SplashPage extends ConsumerWidget {
  const SplashPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final state = ref.watch(splashProvider);

    return Scaffold(
      body: IndicatorOverlay(
        isProcessing: state.loading,
        child: const Center(
          child: TextLogo.large(),
        ),
      ),
    );
  }
}
