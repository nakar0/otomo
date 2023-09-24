import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:otomo/entities/location.dart';

part 'message.freezed.dart';
part 'message.g.dart';

@freezed
class TextMessage with _$TextMessage {
  @JsonSerializable(fieldRename: FieldRename.snake)
  const factory TextMessage({
    required String id,
    String? clientId,
    required String text,
    required Role role,
    required DateTime sentAt,
    required LocationAnalysis locationAnalysis,
  }) = _TextMessage;

  factory TextMessage.fromJson(Map<String, dynamic> json) =>
      _$TextMessageFromJson(json);
}

enum Role {
  user,
  otomo,
}

@freezed
class LocationAnalysis with _$LocationAnalysis {
  @JsonSerializable(fieldRename: FieldRename.snake)
  const factory LocationAnalysis({
    required List<AnalyzedLocation> locations,
    DateTime? analyzedAt,
    String? error,
  }) = _LocationAnalysis;

  factory LocationAnalysis.fromJson(Map<String, dynamic> json) =>
      _$LocationAnalysisFromJson(json);
}

@freezed
class AnalyzedLocation with _$AnalyzedLocation {
  @JsonSerializable(fieldRename: FieldRename.snake)
  const factory AnalyzedLocation({
    required String text,
    required Location location,
  }) = _AnalyzedLocation;

  factory AnalyzedLocation.fromJson(Map<String, dynamic> json) =>
      _$AnalyzedLocationFromJson(json);
}

@freezed
class TextMessageChunk with _$TextMessageChunk {
  const factory TextMessageChunk({
    required String messageId,
    required String text,
    required Role role,
    required DateTime sentAt,
    String? clientId,
    required bool isLast,
  }) = _TextMessageChunk;
}
