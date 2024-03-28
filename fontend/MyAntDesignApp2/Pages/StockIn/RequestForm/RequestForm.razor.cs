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
using Blazored.LocalStorage;
using Microsoft.AspNetCore.Components;
using Microsoft.JSInterop;
using WmsApp.Models;
using WmsApp.Services;

namespace WmsApp.Pages.StockIn
{
    public partial class RequestForm
    {
        private const string buttonSize = "small";
        private string requestFormTitle = "申请入库";

        int _pageIndex = 1;
        int _pageSize = 5;
        StockAddModel stockModel = new();

        List<Warehouse> _warehouses;
        List<Dict> _stockInTypes;
        List<Tenant> _stockInSupplier;
        List<SimplySku> _skuList;

        Table<StockItem> _table;

        Dictionary<int, (bool edit, StockItem data)> editCache = new Dictionary<int, (bool edit, StockItem data)>();
        List<StockItem> skuStockList = new();

        [Inject]
        public IJSRuntime JsRuntime { get; set; }
        [Inject]
        public IBizService BizSrv { get; set; }
        [Inject]
        public ILocalStorageService LocalStorage { get; set; }

        public async Task DownloadTemplate()
        {
            await JsRuntime.InvokeVoidAsync("open", "http://kaw1teau5nc.feishu.cn/docx/JONodSDiaodZQBxE2tnc3kcunkc", "_blank");
        }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();

            DictResp dictResp = await BizSrv.GetDictListAsync();
            WarehouseResp warehouseResp = await BizSrv.GetWarehouseListAsync();
            TenantResp supplierResp = await BizSrv.GetSupplierListAsync();
            SimplySkuResp skuSkuResp = await BizSrv.GetSkuListAsync();

            // 入库类型 = 1
            _stockInTypes = dictResp.Data.Where(a => a.BizType == 1).ToList();
            _warehouses = warehouseResp.Data.ToList();
            _stockInSupplier = supplierResp.Data.ToList();
            _skuList = skuSkuResp.Data.ToList();

        }

        void startEdit(int id)
        {
            (bool edit, StockItem data) editData = editCache[id];
            editCache[id] = (true, editData.data with { }); // add a copy in cache
        }

        void deleteItem(int id)
        {
            int index = skuStockList.FindIndex(item => item.Id == id);
            skuStockList.RemoveAt(index);
            editCache.Remove(id);
        }

        void cancelEdit(int id)
        {
            StockItem data = skuStockList.FirstOrDefault(item => item.Id == id);
            editCache[id] = (false, data); // 恢复 recovery
        }

        void saveEdit(int id)
        {
            int index = skuStockList.FindIndex(item => item.Id == id);
            skuStockList[index] = editCache[id].data; // apply the copy to data source
            editCache[id] = (false, skuStockList[index]); // don't affect rows in editing
        }

        void addItem()
        {
            int lastSkuId = 0;
            if (skuStockList.Count > 0)
            {
                lastSkuId = skuStockList.Last().Id;
            }
            StockItem item = new StockItem
            {
                Id = lastSkuId + 1,
                SkuType = "IIG3000-4G",
                Mac = "00:00:00:00:00:00",
                Iccid = "00000000000000000"
            };
            editCache[item.Id] = (true, item);
            skuStockList.Add(item);
            //
            int _total = skuStockList.Count;
            _pageIndex = _total % _pageSize == 0 ? _total / _pageSize : (_total / _pageSize) + 1;
            // _table.SelectAll();
        }

        // 
        record StockItem
        {
            public int Id { get; set; }
            public string SkuType { get; set; }
            public string Mac { get; set; }
            public string Iccid { get; set; }
        };

    }
}
