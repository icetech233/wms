using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Runtime.Intrinsics.X86;
using System.Threading.Tasks;
using AntDesign;
using Microsoft.AspNetCore.Components;
using MyAntDesignApp2.Models;
using System.Text.Json;
using System.Linq;
using static System.Runtime.InteropServices.JavaScript.JSType;

namespace MyAntDesignApp2.Pages.Settings
{

    public partial class Attrs
    {
        private string _attrPageKey = "规格如下:";
        Attr[] attrList;
        [Inject]
        public HttpClient hc { get; init; }

        [Inject]
        INotificationService _notice { get; init; }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();

            string requestUri = "http://hw.acgzj.cn:8081" +
            "/api/v1/attr/list?s" + +Random.Shared.Next(int.MaxValue);
            var resp = await hc.GetFromJsonAsync<AttrResp>(requestUri);
            Console.WriteLine(resp);
            Console.Out.WriteLine(JsonSerializer.Serialize(resp));
            if (resp.Code != 200)
            {
                // 弹窗 resp.Msg;
                return;
            }
            //resp.Data
            attrList = resp.Data;

            requestUri = "http://hw.acgzj.cn:8081" +
"/api/v1/menu/list?s" + +Random.Shared.Next(int.MaxValue);
            var resp2 = await hc.GetFromJsonAsync<MenuResp>(requestUri);

            Console.WriteLine(resp2);
            Console.Out.WriteLine(JsonSerializer.Serialize(resp2));
        }

        private async Task NoticeWithIcon(string msg)
        {
            await _notice.Open(new NotificationConfig()
            {
                Message = "错误",
                Description = msg,
                NotificationType = NotificationType.Error
            });
        }

        private async void Refresh()
        {
            string requestUri = "http://hw.acgzj.cn:8081" +
"/api/v1/attr/list?s" + Random.Shared.Next(int.MaxValue);
            var resp = hc.GetFromJsonAsync<AttrResp>(requestUri).Result;

            Console.Out.WriteLine(JsonSerializer.Serialize(resp));
            if (resp.Code != 200)
            {
                await NoticeWithIcon(resp.Msg);
                return;
            }
            attrList = resp.Data;
        }


        private void Refresh22()
        {
            var l = attrList.ToList();
            var newId = l.Last().AttrID + 1;
            l.Add(new Attr() { AttrID = newId, AttrName = "test" + newId });
            attrList = l.ToArray();
        }
    }
}
