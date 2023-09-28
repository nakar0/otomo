//
//  Generated code. Do not modify.
//  source: message.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleDescriptor instead')
const Role$json = {
  '1': 'Role',
  '2': [
    {'1': 'UNKNOWN', '2': 0},
    {'1': 'USER', '2': 1},
    {'1': 'OTOMO', '2': 2},
  ],
};

/// Descriptor for `Role`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List roleDescriptor = $convert.base64Decode(
    'CgRSb2xlEgsKB1VOS05PV04QABIICgRVU0VSEAESCQoFT1RPTU8QAg==');

@$core.Deprecated('Use messageDescriptor instead')
const Message$json = {
  '1': 'Message',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'text', '3': 2, '4': 1, '5': 9, '10': 'text'},
    {'1': 'role', '3': 3, '4': 1, '5': 14, '6': '.Role', '10': 'role'},
    {'1': 'sent_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'sentAt'},
    {'1': 'client_id', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.StringValue', '10': 'clientId'},
    {'1': 'location_analysis', '3': 6, '4': 1, '5': 11, '6': '.LocationAnalysis', '10': 'locationAnalysis'},
  ],
};

/// Descriptor for `Message`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageDescriptor = $convert.base64Decode(
    'CgdNZXNzYWdlEg4KAmlkGAEgASgJUgJpZBISCgR0ZXh0GAIgASgJUgR0ZXh0EhkKBHJvbGUYAy'
    'ABKA4yBS5Sb2xlUgRyb2xlEjMKB3NlbnRfYXQYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGlt'
    'ZXN0YW1wUgZzZW50QXQSOQoJY2xpZW50X2lkGAUgASgLMhwuZ29vZ2xlLnByb3RvYnVmLlN0cm'
    'luZ1ZhbHVlUghjbGllbnRJZBI+ChFsb2NhdGlvbl9hbmFseXNpcxgGIAEoCzIRLkxvY2F0aW9u'
    'QW5hbHlzaXNSEGxvY2F0aW9uQW5hbHlzaXM=');

@$core.Deprecated('Use messageChunkDescriptor instead')
const MessageChunk$json = {
  '1': 'MessageChunk',
  '2': [
    {'1': 'message_id', '3': 1, '4': 1, '5': 9, '10': 'messageId'},
    {'1': 'text', '3': 2, '4': 1, '5': 9, '10': 'text'},
    {'1': 'role', '3': 3, '4': 1, '5': 14, '6': '.Role', '10': 'role'},
    {'1': 'sent_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'sentAt'},
    {'1': 'client_id', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.StringValue', '10': 'clientId'},
    {'1': 'is_last', '3': 6, '4': 1, '5': 8, '10': 'isLast'},
  ],
};

/// Descriptor for `MessageChunk`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageChunkDescriptor = $convert.base64Decode(
    'CgxNZXNzYWdlQ2h1bmsSHQoKbWVzc2FnZV9pZBgBIAEoCVIJbWVzc2FnZUlkEhIKBHRleHQYAi'
    'ABKAlSBHRleHQSGQoEcm9sZRgDIAEoDjIFLlJvbGVSBHJvbGUSMwoHc2VudF9hdBgEIAEoCzIa'
    'Lmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSBnNlbnRBdBI5CgljbGllbnRfaWQYBSABKAsyHC'
    '5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWVSCGNsaWVudElkEhcKB2lzX2xhc3QYBiABKAhS'
    'BmlzTGFzdA==');

@$core.Deprecated('Use locationAnalysisDescriptor instead')
const LocationAnalysis$json = {
  '1': 'LocationAnalysis',
  '2': [
    {'1': 'locations', '3': 1, '4': 3, '5': 11, '6': '.AnalyzedLocation', '10': 'locations'},
    {'1': 'analyzed_at', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'analyzedAt'},
    {'1': 'error', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.StringValue', '10': 'error'},
  ],
};

/// Descriptor for `LocationAnalysis`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List locationAnalysisDescriptor = $convert.base64Decode(
    'ChBMb2NhdGlvbkFuYWx5c2lzEi8KCWxvY2F0aW9ucxgBIAMoCzIRLkFuYWx5emVkTG9jYXRpb2'
    '5SCWxvY2F0aW9ucxI7CgthbmFseXplZF9hdBgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1l'
    'c3RhbXBSCmFuYWx5emVkQXQSMgoFZXJyb3IYAyABKAsyHC5nb29nbGUucHJvdG9idWYuU3RyaW'
    '5nVmFsdWVSBWVycm9y');

@$core.Deprecated('Use analyzedLocationDescriptor instead')
const AnalyzedLocation$json = {
  '1': 'AnalyzedLocation',
  '2': [
    {'1': 'text', '3': 1, '4': 1, '5': 9, '10': 'text'},
    {'1': 'location', '3': 2, '4': 1, '5': 11, '6': '.Location', '10': 'location'},
  ],
};

/// Descriptor for `AnalyzedLocation`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List analyzedLocationDescriptor = $convert.base64Decode(
    'ChBBbmFseXplZExvY2F0aW9uEhIKBHRleHQYASABKAlSBHRleHQSJQoIbG9jYXRpb24YAiABKA'
    'syCS5Mb2NhdGlvblIIbG9jYXRpb24=');

@$core.Deprecated('Use monthlySentCountDescriptor instead')
const MonthlySentCount$json = {
  '1': 'MonthlySentCount',
  '2': [
    {'1': 'year_month', '3': 1, '4': 1, '5': 11, '6': '.YearMonth', '10': 'yearMonth'},
    {'1': 'count', '3': 2, '4': 1, '5': 5, '10': 'count'},
    {'1': 'daily_sent_count', '3': 3, '4': 3, '5': 11, '6': '.DailySentCount', '10': 'dailySentCount'},
  ],
};

/// Descriptor for `MonthlySentCount`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List monthlySentCountDescriptor = $convert.base64Decode(
    'ChBNb250aGx5U2VudENvdW50EikKCnllYXJfbW9udGgYASABKAsyCi5ZZWFyTW9udGhSCXllYX'
    'JNb250aBIUCgVjb3VudBgCIAEoBVIFY291bnQSOQoQZGFpbHlfc2VudF9jb3VudBgDIAMoCzIP'
    'LkRhaWx5U2VudENvdW50Ug5kYWlseVNlbnRDb3VudA==');

@$core.Deprecated('Use dailySentCountDescriptor instead')
const DailySentCount$json = {
  '1': 'DailySentCount',
  '2': [
    {'1': 'date', '3': 1, '4': 1, '5': 11, '6': '.Date', '10': 'date'},
    {'1': 'count', '3': 2, '4': 1, '5': 5, '10': 'count'},
  ],
};

/// Descriptor for `DailySentCount`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dailySentCountDescriptor = $convert.base64Decode(
    'Cg5EYWlseVNlbnRDb3VudBIZCgRkYXRlGAEgASgLMgUuRGF0ZVIEZGF0ZRIUCgVjb3VudBgCIA'
    'EoBVIFY291bnQ=');

@$core.Deprecated('Use remainingSendCountDescriptor instead')
const RemainingSendCount$json = {
  '1': 'RemainingSendCount',
  '2': [
    {'1': 'monthly', '3': 1, '4': 1, '5': 11, '6': '.RemainingMonthlySendCount', '10': 'monthly'},
    {'1': 'daily', '3': 2, '4': 1, '5': 11, '6': '.RemainingDailySendCount', '10': 'daily'},
  ],
};

/// Descriptor for `RemainingSendCount`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List remainingSendCountDescriptor = $convert.base64Decode(
    'ChJSZW1haW5pbmdTZW5kQ291bnQSNAoHbW9udGhseRgBIAEoCzIaLlJlbWFpbmluZ01vbnRobH'
    'lTZW5kQ291bnRSB21vbnRobHkSLgoFZGFpbHkYAiABKAsyGC5SZW1haW5pbmdEYWlseVNlbmRD'
    'b3VudFIFZGFpbHk=');

@$core.Deprecated('Use remainingMonthlySendCountDescriptor instead')
const RemainingMonthlySendCount$json = {
  '1': 'RemainingMonthlySendCount',
  '2': [
    {'1': 'year_month', '3': 1, '4': 1, '5': 11, '6': '.YearMonth', '10': 'yearMonth'},
    {'1': 'count', '3': 2, '4': 1, '5': 5, '10': 'count'},
  ],
};

/// Descriptor for `RemainingMonthlySendCount`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List remainingMonthlySendCountDescriptor = $convert.base64Decode(
    'ChlSZW1haW5pbmdNb250aGx5U2VuZENvdW50EikKCnllYXJfbW9udGgYASABKAsyCi5ZZWFyTW'
    '9udGhSCXllYXJNb250aBIUCgVjb3VudBgCIAEoBVIFY291bnQ=');

@$core.Deprecated('Use remainingDailySendCountDescriptor instead')
const RemainingDailySendCount$json = {
  '1': 'RemainingDailySendCount',
  '2': [
    {'1': 'date', '3': 1, '4': 1, '5': 11, '6': '.Date', '10': 'date'},
    {'1': 'count', '3': 2, '4': 1, '5': 5, '10': 'count'},
  ],
};

/// Descriptor for `RemainingDailySendCount`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List remainingDailySendCountDescriptor = $convert.base64Decode(
    'ChdSZW1haW5pbmdEYWlseVNlbmRDb3VudBIZCgRkYXRlGAEgASgLMgUuRGF0ZVIEZGF0ZRIUCg'
    'Vjb3VudBgCIAEoBVIFY291bnQ=');

