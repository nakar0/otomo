import 'package:flutter/material.dart';
import 'package:otomo/views/bases/spaces/spaces.dart';

class HomeWithDraggableScrollableBottomSheet extends StatefulWidget {
  const HomeWithDraggableScrollableBottomSheet({
    super.key,
    required this.maxSheetSize,
    required this.minSheetSize,
    required this.initialSheetSize,
    this.onSheetCreated,
    this.snap = false,
    this.snapSizes,
    this.resizeToAvoidBottomInset,
    this.floatingActionButton,
    this.floatingActionButtonLocation,
    this.behindSheetFloatingActionButton,
    required this.bottomSheetBar,
    required this.bottomSheet,
    required this.child,
  });

  final double maxSheetSize;
  final double minSheetSize;
  final double initialSheetSize;
  final bool snap;
  final List<double>? snapSizes;
  final bool? resizeToAvoidBottomInset;
  final void Function(DraggableScrollableController controller)? onSheetCreated;
  final Widget? floatingActionButton;
  final FloatingActionButtonLocation? floatingActionButtonLocation;
  final Widget? behindSheetFloatingActionButton;
  final Widget bottomSheetBar;
  final Widget bottomSheet;
  final Widget child;

  @override
  State<HomeWithDraggableScrollableBottomSheet> createState() =>
      _HomeWithDraggableScrollableBottomSheetState();
}

class _HomeWithDraggableScrollableBottomSheetState
    extends State<HomeWithDraggableScrollableBottomSheet> {
  static const double _sheetBarHeight = kToolbarHeight;

  final _sheetController = DraggableScrollableController();

  double? _sheetHeight;

  double _initialSheetHeight(BuildContext context) {
    final mediaSize = MediaQuery.of(context).size;
    return mediaSize.height * widget.initialSheetSize;
  }

  @override
  void initState() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      widget.onSheetCreated?.call(_sheetController);
    });
    _sheetController.addListener(() {
      setState(() {
        _sheetHeight = _sheetController.pixels;
      });
    });
    super.initState();
  }

  @override
  void dispose() {
    _sheetController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Scaffold(
      resizeToAvoidBottomInset: widget.resizeToAvoidBottomInset,
      floatingActionButton: widget.floatingActionButton,
      floatingActionButtonLocation: widget.floatingActionButtonLocation,
      body: Stack(
        children: [
          widget.child,
          widget.behindSheetFloatingActionButton ?? Spaces.zero,
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 1),
            child: DraggableScrollableSheet(
              maxChildSize: widget.maxSheetSize,
              initialChildSize: widget.initialSheetSize,
              minChildSize: widget.minSheetSize,
              controller: _sheetController,
              snap: widget.snap,
              snapSizes: widget.snapSizes,
              builder: (context, controller) {
                return SingleChildScrollView(
                  controller: controller,
                  child: Container(
                    decoration: BoxDecoration(
                      color: theme.colorScheme.background,
                      borderRadius: const BorderRadius.vertical(
                        top: Radius.circular(16),
                      ),
                      boxShadow: [
                        BoxShadow(
                          color: theme.appBarTheme.shadowColor!,
                          blurRadius: 4,
                          spreadRadius: 2,
                          offset: const Offset(0, -2),
                        ),
                      ],
                    ),
                    clipBehavior: Clip.hardEdge,
                    height: _sheetHeight ?? _initialSheetHeight(context),
                    child: Column(
                      children: [
                        SizedBox(
                          height: _sheetBarHeight,
                          child: widget.bottomSheetBar,
                        ),
                        Expanded(
                          child: widget.bottomSheet,
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}
