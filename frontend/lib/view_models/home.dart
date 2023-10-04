import 'package:otomo/entities/location.dart';
import 'package:otomo/view_models/boundary/chat.dart';
import 'package:otomo/view_models/chat.dart';
import 'package:otomo/view_models/map.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'home.g.dart';

@riverpod
class Home extends _$Home {
  Stream<TextMessageData> get activatedTextMessageStream =>
      ref.read(chatProvider.notifier).activatedTextMessageStream;
  Stream<Location> get focusedPlaceStream =>
      ref.read(mapProvider.notifier).focusedPlaceStream;

  @override
  void build() {}
}
