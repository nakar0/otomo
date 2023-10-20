import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:otomo/tools/optional.dart';
import 'package:otomo/views/bases/forms/rounded_text_form_field.dart';
import 'package:otomo/views/utils/date_formatter.dart';
import 'package:otomo/views/utils/localizations.dart';

class DateFormField extends StatefulWidget {
  const DateFormField({
    Key? key,
    required this.onConfirmed,
    this.validator,
    this.initialValue,
  }) : super(key: key);

  final ValueChanged<DateTime> onConfirmed;
  final FormFieldValidator<DateTime>? validator;
  final DateTime? initialValue;

  @override
  DateFormFieldState createState() => DateFormFieldState();
}

class DateFormFieldState extends State<DateFormField> {
  final _controller = TextEditingController();

  @override
  void initState() {
    final initialValue = widget.initialValue.apply((v) => _formatDateTime(v));
    if (initialValue != null) _controller.text = initialValue;
    super.initState();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  DateFormat _getFormat(BuildContext context) => DateFormatter.date(context);

  DateTime _parseDateTime(String value) => _getFormat(context).parse(value);
  String _formatDateTime(DateTime dateTime) =>
      _getFormat(context).format(dateTime);

  Future<void> _onTap(BuildContext context) async {
    final pickedDate = await _showDatePicker(context);
    // ignore: always_put_control_body_on_new_line
    if (pickedDate == null) return;
    _controller.text = _formatDateTime(pickedDate);
    widget.onConfirmed(pickedDate);
  }

  Future<DateTime?> _showDatePicker(BuildContext context) async {
    final now = DateTime.now();
    return showDatePicker(
      context: context,
      initialDate: widget.initialValue ?? now,
      firstDate: DateTime(1900),
      lastDate: now,
      cancelText: context.l10n.cancel,
      confirmText: context.l10n.confirm,
      errorInvalidText: context.l10n.inputInvalid(context.l10n.date),
    );
  }

  @override
  Widget build(BuildContext context) {
    return RoundedTextFormField(
      controller: _controller,
      readOnly: true,
      decoration: RoundedTextFormField.defaultDecoration.copyWith(
        suffixIcon: const Icon(Icons.arrow_right_rounded),
      ),
      onTap: () => _onTap(context),
      validator: (value) => widget.validator?.call(
        value.apply((v) => _parseDateTime(v)),
      ),
    );
  }
}
