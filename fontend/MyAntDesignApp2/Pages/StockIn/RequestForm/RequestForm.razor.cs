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
using WmsApp.Models;

namespace WmsApp.Pages.StockIn
{
    public partial class RequestForm
    {
        private string requestFormTitle = "申请入库";
        private StockAddModel stockModel = new();

    }
}
