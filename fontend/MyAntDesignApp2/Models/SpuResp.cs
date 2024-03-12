using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class SpuResp : BaseResp
    {
        public object Debug { get; set; }
        public Spu[] Data { get; set; }
    }

    public class Spu : CommonDate
    {
        public long SpuID { get; set; }
        [DisplayName("编码")]
        public string SpuCode { get; set; }
        public string SpuDesc { get; set; }
    }

    public class SKu : CommonDate
    {
        public long SpuID { get; set; }
        public long SkuID { get; set; }
        [DisplayName("编码")]
        public string SkuCode { get; set; }
        [DisplayName("描述")]
        public string SkuDesc { get; set; }
    }
}
