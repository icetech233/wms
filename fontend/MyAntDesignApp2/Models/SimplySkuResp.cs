using System.Collections.Generic;

namespace WmsApp.Models
{
    public class SimplySkuResp : BaseResp
    {
        public object Debug { get; set; }
        public List<SimplySku> Data { get; set; }
    }

}
