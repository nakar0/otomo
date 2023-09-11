import 'package:flutter/material.dart';
import 'package:otomo/views/cases/home/home_with_draggable_scrollable_bottom_sheet.dart';
import 'package:otomo/views/pages/samples/cases/sample_chat.dart';

class NewHome extends StatefulWidget {
  const NewHome({super.key});

  @override
  State<NewHome> createState() => _NewHomeState();
}

class _NewHomeState extends State<NewHome> {
  DraggableScrollableController? _sheetController;

  void onSheetCreated(DraggableScrollableController controller) {
    _sheetController = controller;
  }

  @override
  Widget build(BuildContext context) {
    return HomeWithDraggableScrollableBottomSheet(
      initialSheetSize: 0.1,
      minSheetSize: 0.0,
      maxSheetSize: 0.9,
      snap: true,
      snapSizes: const [0.3, 0.6],
      onSheetCreated: onSheetCreated,
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          _sheetController?.animateTo(
            0.6,
            curve: Curves.ease,
            duration: const Duration(milliseconds: 500),
          );
        },
        child: const Icon(Icons.chat),
      ),
      bottomSheetBar: Container(color: Colors.red),
      bottomSheet: const SampleChat(),
      child: const Placeholder(),
    );
  }
}
