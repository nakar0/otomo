// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'message.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_Message _$$_MessageFromJson(Map<String, dynamic> json) => _$_Message(
      id: json['id'] as String,
      text: json['text'] as String,
      role: $enumDecode(_$RoleEnumMap, json['role']),
      sentAt: DateTime.parse(json['sentAt'] as String),
    );

Map<String, dynamic> _$$_MessageToJson(_$_Message instance) =>
    <String, dynamic>{
      'id': instance.id,
      'text': instance.text,
      'role': _$RoleEnumMap[instance.role]!,
      'sentAt': instance.sentAt.toIso8601String(),
    };

const _$RoleEnumMap = {
  Role.user: 'user',
  Role.otomo: 'otomo',
};
