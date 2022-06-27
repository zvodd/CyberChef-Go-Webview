mkdir libs/webview2

curl -sSL "https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2/1.0.1245.22" | tar -xf - -C libs/webview2

cp libs/webview2/build/native/x64/WebView2Loader.dll build

export CGO_CXXFLAGS="-I$(pwd)/libs/webview2/build/native/include"
export CGO_LDFLAGS="-L$(pwd)/libs/webview2/build/native/x64"

#go build -ldflags="-H windowsgui" -o build/cv.exe main.go