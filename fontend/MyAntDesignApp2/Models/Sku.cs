using System.ComponentModel;

namespace WmsApp.Models
{
    public class Sku : CommonDate
    {
        public long SpuID { get; set; }
        public long SkuID { get; set; }
        [DisplayName("编码")]
        public string SkuCode { get; set; }
        [DisplayName("描述")]
        public string SkuDesc { get; set; }
    }

}
