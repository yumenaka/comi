package components

import (
	"fmt"
	"github.com/ebitenui/ebitenui"
	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yumenaka/comigo/application/ebitenui/model"
	"github.com/yumenaka/comigo/application/ebitenui/resources"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
)

func HeaderContainer(res *resources.UIResources, readerConfig *model.ReaderConfig, ui *ebitenui.UI) widget.PreferredSizeLocateableWidget {

	// 设置文本颜色
	textColor := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	// 加载按钮文字字体
	face, _ := loadFont(20)
	// headerContainer 是一个新的容器，用于包含标题文本和按钮。
	headerContainer := widget.NewContainer(
		// header容器的背景颜色
		widget.ContainerOpts.BackgroundImage(eimage.NewNineSliceColor(color.NRGBA{R: 245, G: 245, B: 228, A: 255})),
		// 设置容器的布局
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			// 使用 Columns 参数来定义列的数量。
			widget.GridLayoutOpts.Columns(7),
			// 使用 ColumnStretch 和 RowStretch 参数来分别定义列和行的拉伸因子。
			// 只支持布尔值，true表示拉伸，false表示不拉伸。
			widget.GridLayoutOpts.Stretch([]bool{false, false, false, true, false, false, false}, []bool{true}),
			//内边距的大小
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(8)),
			//网格布局的间距，c 列间距，r行间距。
			widget.GridLayoutOpts.Spacing(8, 8),
		)),
		// 统一设定布局设定
		widget.ContainerOpts.WidgetOpts(
			//widget.WidgetOpts.MinSize(80, 20),
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				MaxHeight:          56,
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition:   widget.GridLayoutPositionCenter,
			}),
		),
	)
	// 服务器设置按钮
	serverButton := widget.NewButton(
		// 指定要使用的图像
		widget.ButtonOpts.Image(res.Button.Image),
		// 指定按钮的文本、字体和颜色
		widget.ButtonOpts.Text(fmt.Sprintf("Server"), res.Button.Face, res.Button.Text),
		// 按钮文本填充
		widget.ButtonOpts.TextPadding(res.Button.Padding),
		// 添加一个处理程序以响应点击按钮事件
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println(fmt.Sprintf("Server Button Clicked!"))
		}),
		widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Entered: " + args.Button.Text().Label) }),
		widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Exited: " + args.Button.Text().Label) }),
		//// 单独某个按钮的布局设定
		//widget.ButtonOpts.WidgetOpts(
		//	widget.WidgetOpts.MinSize(80, 20),
		//	widget.WidgetOpts.LayoutData(widget.GridLayoutData{
		//		MaxHeight:          40,
		//		HorizontalPosition: widget.GridLayoutPositionCenter,
		//		VerticalPosition:   widget.GridLayoutPositionCenter,
		//	}),
		//),
	)
	headerContainer.AddChild(serverButton)

	// Upload按钮
	uploadButton := widget.NewButton(
		// 指定要使用的图像
		widget.ButtonOpts.Image(res.Button.Image),
		// 指定按钮的文本、字体和颜色
		widget.ButtonOpts.Text(fmt.Sprintf("Upload"), face, &widget.ButtonTextColor{
			Idle: color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff},
		}),
		// 按钮文本填充
		widget.ButtonOpts.TextPadding(res.Button.Padding),
		// 添加一个处理程序以响应点击按钮事件
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println(fmt.Sprintf("Upload Button Clicked!"))
		}),
		// 设置按钮的通用选项
		widget.ButtonOpts.WidgetOpts(
			// 布局设置，将按钮水平和垂直居中
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	headerContainer.AddChild(uploadButton)

	// add the button as a child of the container
	headerContainer.AddChild(SortButton(res))

	// 一个新的文本小部件，用于显示文本。
	titleText := widget.NewText(
		widget.TextOpts.Text(readerConfig.Title, res.Button.Face, textColor),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		//要配置单个小部件与其兄弟小部件有不同的布局，可以在小部件上设置一个可选的“布局数据”。
		//布局数据的类型取决于所使用的布局实现。例如，RowLayout 需要使用 RowLayoutData。
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{}),
		),
	)
	headerContainer.AddChild(titleText)
	QRCodeWindow := QRCodeWindow()
	// QRCode按钮
	qrcodeButton := widget.NewButton(
		// 指定要使用的图像
		widget.ButtonOpts.Image(res.Button.Image),
		// 指定按钮的文本、字体和颜色
		widget.ButtonOpts.Text(fmt.Sprintf("QRCode"), face, &widget.ButtonTextColor{
			Idle: color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff},
		}),
		// 按钮文本填充
		widget.ButtonOpts.TextPadding(res.Button.Padding),
		// 添加一个处理程序以响应点击按钮事件
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// 显示QRCode窗口
			println(fmt.Sprintf("QRCode Button Clicked!"))
			//获取内容的首选大小
			x, y := QRCodeWindow.Contents.PreferredSize()
			//创建一个具有内容首选大小的矩形
			r := image.Rect(0, 0, x, y)
			//如果窗口全屏
			if readerConfig.WindowFullScreen {
				w, h := ebiten.Monitor().Size()
				//使用 Add 方法将窗口移动到指定点
				r = r.Add(image.Point{X: (w / 2) - 200, Y: (h / 2) - 150})
			} else {
				//使用 Add 方法将窗口移动到指定点
				r = r.Add(image.Point{X: (readerConfig.Width / 2) - 200, Y: (readerConfig.Height / 2) - 150})
			}
			//将窗口位置设置到矩形。
			QRCodeWindow.SetLocation(r)
			//将窗口添加到用户界面。
			//注意：如果窗口已经添加，这将只移动窗口，而不会添加重复项。
			ui.AddWindow(QRCodeWindow)
		}),
		// 设置按钮的通用选项
		widget.ButtonOpts.WidgetOpts(
			// 布局设置，将按钮水平和垂直居中
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	headerContainer.AddChild(qrcodeButton)

	// FullScreen按钮
	fullScreenButton := widget.NewButton(
		// 指定要使用的图像
		widget.ButtonOpts.Image(res.Button.Image),
		// 指定按钮的文本、字体和颜色
		widget.ButtonOpts.Text(fmt.Sprintf("FullScreen"), face, &widget.ButtonTextColor{
			Idle: color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff},
		}),
		// 按钮文本填充
		widget.ButtonOpts.TextPadding(res.Button.Padding),
		// 添加一个处理程序以响应点击按钮事件
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println(fmt.Sprintf("FullScreen Button Clicked!"))
			readerConfig.WindowFullScreen = !readerConfig.WindowFullScreen
			ebiten.SetFullscreen(readerConfig.WindowFullScreen)
		}),
		// 设置按钮的通用选项
		widget.ButtonOpts.WidgetOpts(
			// 布局设置，将按钮水平和垂直居中
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	headerContainer.AddChild(fullScreenButton)

	// 设置按钮
	settingButton := widget.NewButton(
		// 指定要使用的图像
		widget.ButtonOpts.Image(res.Button.Image),
		// 指定按钮的文本、字体和颜色
		widget.ButtonOpts.Text(fmt.Sprintf("Setting"), face, &widget.ButtonTextColor{
			Idle: color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff},
		}),
		// 按钮文本填充
		widget.ButtonOpts.TextPadding(res.Button.Padding),
		// 添加一个处理程序以响应点击按钮事件
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println(fmt.Sprintf("Setting Button Clicked!"))
		}),
		// 设置按钮的通用选项
		widget.ButtonOpts.WidgetOpts(
			// 锚点布局设置，将按钮水平和垂直居中
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	headerContainer.AddChild(settingButton)
	return headerContainer
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := eimage.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := eimage.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := eimage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
