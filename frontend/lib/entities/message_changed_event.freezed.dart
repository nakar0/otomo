// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'message_changed_event.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$TextMessageChangedEvent {
  String get messageId => throw _privateConstructorUsedError;
  ChangedEventType get type => throw _privateConstructorUsedError;
  TextMessage? get data => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $TextMessageChangedEventCopyWith<TextMessageChangedEvent> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TextMessageChangedEventCopyWith<$Res> {
  factory $TextMessageChangedEventCopyWith(TextMessageChangedEvent value,
          $Res Function(TextMessageChangedEvent) then) =
      _$TextMessageChangedEventCopyWithImpl<$Res, TextMessageChangedEvent>;
  @useResult
  $Res call({String messageId, ChangedEventType type, TextMessage? data});

  $TextMessageCopyWith<$Res>? get data;
}

/// @nodoc
class _$TextMessageChangedEventCopyWithImpl<$Res,
        $Val extends TextMessageChangedEvent>
    implements $TextMessageChangedEventCopyWith<$Res> {
  _$TextMessageChangedEventCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? messageId = null,
    Object? type = null,
    Object? data = freezed,
  }) {
    return _then(_value.copyWith(
      messageId: null == messageId
          ? _value.messageId
          : messageId // ignore: cast_nullable_to_non_nullable
              as String,
      type: null == type
          ? _value.type
          : type // ignore: cast_nullable_to_non_nullable
              as ChangedEventType,
      data: freezed == data
          ? _value.data
          : data // ignore: cast_nullable_to_non_nullable
              as TextMessage?,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $TextMessageCopyWith<$Res>? get data {
    if (_value.data == null) {
      return null;
    }

    return $TextMessageCopyWith<$Res>(_value.data!, (value) {
      return _then(_value.copyWith(data: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$_TextMessageChangedEventCopyWith<$Res>
    implements $TextMessageChangedEventCopyWith<$Res> {
  factory _$$_TextMessageChangedEventCopyWith(_$_TextMessageChangedEvent value,
          $Res Function(_$_TextMessageChangedEvent) then) =
      __$$_TextMessageChangedEventCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String messageId, ChangedEventType type, TextMessage? data});

  @override
  $TextMessageCopyWith<$Res>? get data;
}

/// @nodoc
class __$$_TextMessageChangedEventCopyWithImpl<$Res>
    extends _$TextMessageChangedEventCopyWithImpl<$Res,
        _$_TextMessageChangedEvent>
    implements _$$_TextMessageChangedEventCopyWith<$Res> {
  __$$_TextMessageChangedEventCopyWithImpl(_$_TextMessageChangedEvent _value,
      $Res Function(_$_TextMessageChangedEvent) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? messageId = null,
    Object? type = null,
    Object? data = freezed,
  }) {
    return _then(_$_TextMessageChangedEvent(
      messageId: null == messageId
          ? _value.messageId
          : messageId // ignore: cast_nullable_to_non_nullable
              as String,
      type: null == type
          ? _value.type
          : type // ignore: cast_nullable_to_non_nullable
              as ChangedEventType,
      data: freezed == data
          ? _value.data
          : data // ignore: cast_nullable_to_non_nullable
              as TextMessage?,
    ));
  }
}

/// @nodoc

class _$_TextMessageChangedEvent implements _TextMessageChangedEvent {
  const _$_TextMessageChangedEvent(
      {required this.messageId, required this.type, this.data});

  @override
  final String messageId;
  @override
  final ChangedEventType type;
  @override
  final TextMessage? data;

  @override
  String toString() {
    return 'TextMessageChangedEvent(messageId: $messageId, type: $type, data: $data)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_TextMessageChangedEvent &&
            (identical(other.messageId, messageId) ||
                other.messageId == messageId) &&
            (identical(other.type, type) || other.type == type) &&
            (identical(other.data, data) || other.data == data));
  }

  @override
  int get hashCode => Object.hash(runtimeType, messageId, type, data);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_TextMessageChangedEventCopyWith<_$_TextMessageChangedEvent>
      get copyWith =>
          __$$_TextMessageChangedEventCopyWithImpl<_$_TextMessageChangedEvent>(
              this, _$identity);
}

abstract class _TextMessageChangedEvent implements TextMessageChangedEvent {
  const factory _TextMessageChangedEvent(
      {required final String messageId,
      required final ChangedEventType type,
      final TextMessage? data}) = _$_TextMessageChangedEvent;

  @override
  String get messageId;
  @override
  ChangedEventType get type;
  @override
  TextMessage? get data;
  @override
  @JsonKey(ignore: true)
  _$$_TextMessageChangedEventCopyWith<_$_TextMessageChangedEvent>
      get copyWith => throw _privateConstructorUsedError;
}
