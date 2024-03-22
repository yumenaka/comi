import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../models/book.dart';
import '../models/remote_server.dart';

// 这个Widget是Home页面的根部件。
class ScrollMode extends StatefulWidget {
  const ScrollMode({super.key, required this.title});
  // 这个小部件是应用的主页。它是有状态的，意味着它有一个包含影响其外观的字段的 State 对象（在下面定义）。
  // 这个类是状态的配置。它保存了父级（在这种情况下是 App 小部件）提供的值（在这种情况下是标题），并由状态的 build 方法使用。
  // Widget 子类中的字段始终标记为 "final"。
  final String title;

  @override
  State<ScrollMode> createState() => _ScrollModeState();
}

class _ScrollModeState extends State<ScrollMode> {

  @override
  void initState() {
    super.initState();
  }

  /// 获取书籍
  Future<Book> initBooks() async {
    Future<Book>? book = getBook(); // 调用函数并初始化参数
    return book.then((value) => value);
  }

  @override
  Widget build(BuildContext context) {
    //  画面の高さを取得する
    final mediaQueryData = MediaQuery.of(context);
    final headerHeight = mediaQueryData.size.height * 0.06;

    // 异步UI更新（FutureBuilder、StreamBuilder）
    // https://book.flutterchina.club/chapter7/futurebuilder_and_streambuilder.html
    Widget booksWidget = FutureBuilder<Book>(
      future: getBook(),
      initialData: null,
      // snapshot会包含当前异步任务的状态信息及结果信息
      builder: (context, snapshot) {
        // 请求结束
        if (snapshot.hasData) {
          debugPrint(snapshot.data!.title.toString());
          return ListView.builder(
            itemCount: snapshot.data!.pages!.length,
            itemBuilder: (context, index) {
              return Container(
                color: Colors.white,
                padding: const EdgeInsets.all(2.0),
                margin: const EdgeInsets.fromLTRB(8.0, 8.0, 8.0, 8.0),
                child: ListTile(
                  //snapshot.data!.pages!.length
                    trailing: Image.network("${Provider.of<RemoteServer>(context).remoteHost}/${snapshot.data!.pages![index].url}"),
                    onTap: () {
                      debugPrint("remoteHost:${Provider.of<RemoteServer>(context, listen: false).remoteHost}");
                    }
                ),
              );
            },
          );
        } else if (snapshot.hasError) {
          return Text('${snapshot.error}');
        }
        // 默认情况下，显示一个进度条
        return const CircularProgressIndicator();
      },
    );

    // 每次调用setState时，此方法都会重新运行，例如上面的_incrementCounter方法。
    // Flutter框架已经进行了优化，使重新运行build方法变得快速，因此您只需重新构建需要更新的内容，
    // 而不必逐个更改小部件的实例。
    return Scaffold(
      backgroundColor: Colors.greenAccent[100],
      appBar: AppBar(
        title: Text(widget.title,
            style: Theme.of(context).textTheme.headlineMedium),
        backgroundColor: Colors.lightBlue[200], ////appbarの背景色を設定する
        toolbarHeight: headerHeight, //appbarの高さを設定する
        centerTitle: true, //タイトルを中央に配置
      ),
      body: booksWidget,
      bottomNavigationBar: BottomNavigationBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.store),
            label: 'Store',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: 'Settings',
          ),
        ],
        selectedItemColor: Colors.amber[800],
        onTap: (int index) {
          debugPrint("index:$index");
        },
      ),
    );
  }
}
