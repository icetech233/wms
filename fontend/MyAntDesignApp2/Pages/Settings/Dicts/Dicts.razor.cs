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
using Microsoft.JSInterop;
using WmsApp.Models;
using WmsApp.Services;
using Blazored.LocalStorage;

namespace WmsApp.Pages.Settings
{
    public partial class Dicts
    {
        const string _PageKey = "字典选项:";
        string size = "middle";
        bool _visible = false;
        string modalTitle = "添加仓库"; // 编辑仓库

        [Inject]
        public IDictService DictSrv { get; set; }
        [Inject]
        public ILocalStorageService LocalStorage { get; set; }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();
            await Refresh();
        }

        private async Task Refresh()
        {
            DictResp resp = await DictSrv.GetDictListAsync();
            Console.WriteLine($"refresh code {resp.Code} count {resp.Data?.Length}");
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
            }
            //
            dictList = resp.Data;
            await LocalStorage.SetItemAsync("dict_context", dictList);
        }

        private void Delete(long id)
        {
            // 请求网络 delete 接口
            // warehouseList = warehouseList.Where(x => x.WarehouseID != id).ToArray();
        }

    }
}
