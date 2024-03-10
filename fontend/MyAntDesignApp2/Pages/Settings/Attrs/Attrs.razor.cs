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
using Microsoft.Extensions.Logging;
// Microsoft.Extensions.Logging;

namespace MyAntDesignApp2.Pages.Settings
{

    public partial class Attrs
    {
        [Inject]
        private ILogger<Attrs> logger { get; set; }

        string modalTitle = "添加规格"; // 编辑规格
        string size = "middle";
        bool _visible = false;
        readonly string _attrPageKey = "规格如下:";

        public Attr[] attrList;
        [Inject]
        public HttpClient hc { get; init; }

        protected override async Task OnInitializedAsync()
        {
            logger.LogWarning("Someone has clicked me!");
            // OnAfterRenderAsync
            await base.OnInitializedAsync();

            string requestUri = "http://hw.acgzj.cn:8081" +
            "/api/v1/attr/list?s=" + Random.Shared.Next(int.MaxValue);
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
"/api/v1/menu/list?s=" + Random.Shared.Next(int.MaxValue);
            var resp2 = await hc.GetFromJsonAsync<MenuResp>(requestUri);

            Console.WriteLine(resp2);
            Console.WriteLine(JsonSerializer.Serialize(resp2));
        }

        private async Task Refresh()
        {
            string requestUri = "http://hw.acgzj.cn:8081" +
"/api/v1/attr/list?s=" + Random.Shared.Next(int.MaxValue);
            AttrResp resp = await hc.GetFromJsonAsync<AttrResp>(requestUri);

            // Console.WriteLine(JsonSerializer.Serialize(resp));
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
            }
            attrList = resp.Data;

        }

    }
}
