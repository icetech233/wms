using System.Collections.Generic;
using System.ComponentModel;

namespace WmsApp.Models
{
    public class Spu : CommonDate
    {
        public long SpuID { get; set; }
        [DisplayName("编码")]
        public string SpuCode { get; set; }
        public string SpuDesc { get; set; }
        public List<Sku> SkuList { get; set; }
    }

}
