using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json;
using System.Threading.Tasks;
using AntDesign;
using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Logging;
using MyAntDesignApp2.Models;

namespace MyAntDesignApp2.Pages.Settings
{
    public partial class Dicts
    {
        const string _PageKey = "字典选项:";
        string size = "middle";
        bool _visible = false;
        string modalTitle = "添加仓库"; // 编辑仓库

        [Inject]
        public HttpClient hc { get; set; }

        protected override async Task OnInitializedAsync()
        {
            // logger.LogWarning("OnInitializedAsync has clicked me!");
            await base.OnInitializedAsync();

            await Refresh();
        }

        private DictModel model = new();
        public class DictModel
        {
            public long DictID { get; set; }
            public int BizType { get; set; }
            public string DictName { get; set; }
            public string DictValue { get; set; }
            [Required]
            public string Address { get; set; }
        }

        private async Task Refresh()
        {
            string requestUri = "/api/v1/dict/list?s=" + Random.Shared.Next(int.MaxValue);
            DictResp resp = await hc.GetFromJsonAsync<DictResp>(requestUri);
            Console.WriteLine($"refresh code {resp.Code} count {resp.Data?.Length}");
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
            }
            dictList = resp.Data;
        }

        private void Delete(long id)
        {
            // 请求网络 delete 接口
            // warehouseList = warehouseList.Where(x => x.WarehouseID != id).ToArray();
        }

    }
}
