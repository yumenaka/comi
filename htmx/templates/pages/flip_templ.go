// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/yumenaka/comi/entity"
	"github.com/yumenaka/comi/htmx/state"
	"github.com/yumenaka/comi/htmx/templates/components"
)

func FlipScripts() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_FlipScripts_0f6e`,
		Function: `function __templ_FlipScripts_0f6e(){//滚动到顶部
function FlipToTop(FlipDuration) {
    let FlipStep = -window.FlipY / (FlipDuration / 15),
        FlipInterval = setInterval(function () {
            if (window.FlipY !== 0) {
                window.FlipBy(0, FlipStep);
            }
            else clearInterval(FlipInterval);
        }, 15);
}
// Button ID为BackTopButton的元素，点击后滚动到顶部
document.getElementById("BackTopButton").addEventListener("click", function () {
    FlipToTop(500);
});

//滚动到一定位置显示返回顶部按钮
let FlipTopSave = 0
let FlipDownFlag = false
let step = 0
function onFlip() {
    let FlipTop = document.documentElement.FlipTop || document.body.FlipTop;
    FlipDownFlag = FlipTop > FlipTopSave;
    //防手抖,小于一定数值状态就不变 Math.abs()会导致报错
    step = FlipTopSave - FlipTop;
    // console.log("this.FlipDownFlag:",this.FlipDownFlag,"FlipTop:",FlipTop,"step:", step);
    FlipTopSave = FlipTop
    if (step < -10 || step > 10) {
        showBackTopFlag = ((FlipTop > 400) && !FlipDownFlag);
        if (showBackTopFlag) {
            document.getElementById("BackTopButton").style.display = "block";
        } else {
            document.getElementById("BackTopButton").style.display = "none";
        }
    }
}
window.addEventListener("Flip", onFlip);

//可见区域变化的时候改变页面状态
function onResize() {
    this.FlipModeConfig.imageMaxWidth = window.innerWidth
    this.clientWidth = document.documentElement.clientWidth
    this.clientHeight = document.documentElement.clientHeight
    // var aspectRatio = window.innerWidth / window.innerHeight
    this.aspectRatio = this.clientWidth / this.clientHeight
    // 为了调试的时候方便,阈值是正方形
    if (this.aspectRatio > (19 / 19)) {
        this.FlipModeConfig.isLandscapeMode = true
        this.FlipModeConfig.isPortraitMode = false
    } else {
        this.FlipModeConfig.isLandscapeMode = false
        this.FlipModeConfig.isPortraitMode = true
    }
}
//文档视图调整大小时触发 resize 事件。 https://developer.mozilla.org/zh-CN/docs/Web/API/Window/resize_event
window.addEventListener("resize", this.onResize);

//获取鼠标位置,决定是否打开设置面板
function onMouseClick(e) {
    this.clickX = e.x //获取鼠标的X坐标（鼠标与屏幕左侧的距离,单位为px）
    this.clickY = e.y //获取鼠标的Y坐标（鼠标与屏幕顶部的距离,单位为px）
    //浏览器的视口,不包括工具栏和滚动条:
    let innerWidth = window.innerWidth
    let innerHeight = window.innerHeight
    //设置区域为正方形，边长按照宽或高里面，比较小的值决定
    const setArea = 0.15;
    // innerWidth >= innerHeight 的情况下
    let MinY = innerHeight * (0.5 - setArea);
    let MaxY = innerHeight * (0.5 + setArea);
    let MinX = innerWidth * 0.5 - (MaxY - MinY) * 0.5;
    let MaxX = innerWidth * 0.5 + (MaxY - MinY) * 0.5;
    if (innerWidth < innerHeight) {
        MinX = innerWidth * (0.5 - setArea);
        MaxX = innerWidth * (0.5 + setArea);
        MinY = innerHeight * 0.5 - (MaxX - MinX) * 0.5;
        MaxY = innerHeight * 0.5 + (MaxX - MinX) * 0.5;
    }
    //在设置区域
    let inSetArea = false
    if ((this.clickX > MinX && this.clickX < MaxX) && (this.clickY > MinY && this.clickY < MaxY)) {
        console.log("点中了设置区域！");
        inSetArea = true
    }
    if (inSetArea) {
        //获取ID为 OpenSettingButton的元素，然后模拟点击
		document.getElementById("OpenSettingButton").click();
    }
}
//获取鼠标位置,决定是否打开设置面板
function onMouseMove(e) {
    this.clickX = e.x //获取鼠标的X坐标（鼠标与屏幕左侧的距离,单位为px）
    this.clickY = e.y //获取鼠标的Y坐标（鼠标与屏幕顶部的距离,单位为px）
    //浏览器的视口,不包括工具栏和滚动条:
    let innerWidth = window.innerWidth
    let innerHeight = window.innerHeight
    //设置区域为正方形，边长按照宽或高里面，比较小的值决定
    const setArea = 0.15;
    // innerWidth >= innerHeight 的情况下
    let MinY = innerHeight * (0.5 - setArea);
    let MaxY = innerHeight * (0.5 + setArea);
    let MinX = innerWidth * 0.5 - (MaxY - MinY) * 0.5;
    let MaxX = innerWidth * 0.5 + (MaxY - MinY) * 0.5;
    if (innerWidth < innerHeight) {
        MinX = innerWidth * (0.5 - setArea);
        MaxX = innerWidth * (0.5 + setArea);
        MinY = innerHeight * 0.5 - (MaxX - MinX) * 0.5;
        MaxY = innerHeight * 0.5 + (MaxX - MinX) * 0.5;
    }
    //在设置区域
    let inSetArea = false
    if ((this.clickX > MinX && this.clickX < MaxX) && (this.clickY > MinY && this.clickY < MaxY)) {
        inSetArea = true
    }
    if (inSetArea) {
        //console.log("在设置区域！");
        e.currentTarget.style.cursor = 'url(/static/images/SettingsOutline.png), pointer';
    } else {
        e.currentTarget.style.cursor = '';
    }
}
//获取ID为 mouseMoveArea 的元素
let mouseMoveArea = document.getElementById("mouseMoveArea")
// 鼠标移动的时候触发移动事件
mouseMoveArea.addEventListener('mousemove', onMouseMove)
// 点击的时候触发点击事件
mouseMoveArea.addEventListener('click', onMouseClick)
// 触摸的时候也触发点击事件
mouseMoveArea.addEventListener('touchstart', onMouseClick)
}`,
		Call:       templ.SafeScript(`__templ_FlipScripts_0f6e`),
		CallInline: templ.SafeScriptInline(`__templ_FlipScripts_0f6e`),
	}
}

func FlipMainArea(s *state.GlobalState, book *entity.Book) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"mouseMoveArea\" class=\"flex flex-col items-center justify-center flex-1 w-full max-w-full bg-base-100 text-base-content\" :class=\"(theme.toString() ===&#39;light&#39;||theme.toString() ===&#39;dark&#39;||theme.toString() ===&#39;retro&#39;||theme.toString() ===&#39;lofi&#39;||theme.toString() ===&#39;nord&#39;) &amp;&amp; &#39;bg-base-300&#39;\"><div class=\"manga_area\" id=\"MangaMain\" @click.stop=\"onMouseClick\" @mousemove.stop=\"onMouseMove\" @mouseleave.stop=\"onMouseLeave\"><div class=\"manga_area_img_div\"><!-- 非自动拼合模式最简单,直接显示一张图 --><img class=\"w-auto h-auto\" v-bind:src=\"imageParametersString(book.pages.images[nowPageNum - 1].url)\n            \" v-bind:alt=\"nowPageNum.toString()\"><!-- 简单拼合双页,不管单双页什么的 --><img v-if=\"!FlipModeConfig.autoDoublePageModeFlag &amp;&amp;\n              FlipModeConfig.doublePageModeFlag &amp;&amp;\n              nowPageNum &lt; book.page_count\n            \" v-bind:src=\"imageParametersString(book.pages.images[nowPageNum].url)\n            \" v-bind:alt=\"(nowPageNum + 1).toString()\"><!-- 自动拼合模式当前页,如果开启自动拼合,右边可能显示拼合页 --><img v-if=\"FlipModeConfig.autoDoublePageModeFlag &amp;&amp;\n              nowPageNum &lt; book.page_count &amp;&amp;\n            nowAndNextPageIsSingle()\n            \" v-bind:src=\"imageParametersString(book.pages.images[nowPageNum].url)\n            \" v-bind:alt=\"(nowPageNum + 1).toString()\"></div></div><noscript class=\"loading-lazy\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for key, image := range book.Pages.Images {
			if key <= 2 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img class=\"m-2 max-w-full lg:max-w-[800px] rounded shadow-lg\" src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(image.Url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/flip.templ`, Line: 170, Col: 83}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if key > 2 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img loading=\"lazy\" class=\"m-2 max-w-full lg:max-w-[800px] rounded shadow-lg\" src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(image.Url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/flip.templ`, Line: 173, Col: 98}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</noscript></div><button id=\"BackTopButton\" style=\"display: none\" class=\"fixed bottom-4 right-4 bg-blue-500 text-white rounded-full w-10 h-10 flex items-center justify-center shadow-lg\"><svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 512 512\"><path d=\"M256 48C141.13 48 48 141.13 48 256s93.13 208 208 208s208-93.13 208-208S370.87 48 256 48zm96 270.63l-96-96l-96 96L137.37 296L256 177.37L374.63 296z\" fill=\"currentColor\"></path></svg></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = FlipScripts().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FlipDrawerSlot() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// FlipPage 定义 BodyHTML
func FlipPage(s *state.GlobalState, book *entity.Book) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = components.Header(components.HeaderProps{
			Title:           getPageTitle(s.RequestBookID),
			ShowReturnIcon:  true,
			ReturnUrl:       getReturnUrl(s.RequestBookID),
			SetDownLoadLink: false,
			InShelf:         false,
			DownLoadLink:    "",
			SetTheme:        true,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = FlipMainArea(s, book).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Footer(s.Version).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Drawer(s, FlipDrawerSlot()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.QRCode(s).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
