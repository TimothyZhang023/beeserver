<!DOCTYPE html>

<html>
<head>
  <title>MiPushFramework Release Page</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {

      text-align: center;
      font-size: 42px;
      padding: 0px 0 0px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 80px 0  40px 0;
    }

    article {
      line-height: 1.8;
      text-align: center;
      padding: 10px 0;
      color: #999;
       margin: auto;
     }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
   .tg  { min-width: 800px;}

.tg  {border-collapse:collapse;border-spacing:0;border-color:#aaa;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:0px;overflow:hidden;word-break:normal;border-top-width:1px;border-bottom-width:1px;border-color:#aaa;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:0px;overflow:hidden;word-break:normal;border-top-width:1px;border-bottom-width:1px;border-color:#aaa;color:#fff;background-color:#f38630;}
.tg .tg-0lax{text-align:left;vertical-align:top}
.tg .tg-0pky{border-color:inherit;text-align:left;vertical-align:top}


  </style>
</head>

<body>
  <header>
    <h1 class="logo">MiPushFramework Release Page</h1>
    <div class="description">
     </div>
  </header>

      <article>
      <center class="main" >
       <table class="tg">
        <tr>
          <th class="tg-0pky">版本</th>
          <th class="tg-0pky">UI界面</th>
          <th class="tg-0pky">PUSH服务</th>
        </tr>

         {{range $i, $releaseItem := .releases}}

        <tr>
          <td class="tg-0pky">{{$releaseItem.Version}}</td>
          <td class="tg-0pky"><a href="{{$releaseItem.GetManagerUrlCN}}" >国内(HTTP)</a> &nbsp; &nbsp;<a href="{{$releaseItem.GetManagerUrl}}" >备用(HTTPS)</a>  </td>
          <td class="tg-0pky"><a href="{{$releaseItem.GetXmsfUrlCN}}" >国内(HTTP)</a> &nbsp;&nbsp; <a href="{{$releaseItem.GetXmsfUrl}}" >备用(HTTPS)</a>  </td>
        </tr>
            {{end}}

      </table>



      </center>
      </article>

  <footer>
    <div class="author">
      Official website:
      <a href="https://github.com/Trumeet/MiPushFramework">https://github.com/Trumeet/MiPushFramework</a> /
      Contact us:
      <a class="tg" href="https://t.me/mipushframework">Telegram Group</a>
    </div>
  </footer>
  <div class="backdrop"></div>

 </body>
</html>
