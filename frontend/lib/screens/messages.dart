import 'package:flutter/material.dart';
import '../models/message.dart';  // 假设你已经定义了Message模型

class Messages extends StatelessWidget {
  final List<Message> messages = [];

  Messages({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('消息'),
      ),
      body: ListView.builder(
        itemCount: messages.length,
        itemBuilder: (context, index) {
          final msg = messages[index];
          return ListTile(
            leading: const Icon(Icons.message),
            title: Text(msg.content),
            subtitle: Text(msg.timestamp.toString()),
          );
        },
      ),
    );
  }
}
