using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class Warehouse : CommonDate
    {
        [DisplayName("ID")]
        public long WarehouseID { get; set; }
        [DisplayName("启用")]
        public bool Enable { get; set; }
        [DisplayName("编码")]
        public string WarehouseCode { get; set; }
        [DisplayName("名称")]
        public string WarehouseName { get; set; }
        [DisplayName("仓库管理员")]
        public string Supervisor { get; set; }
        [DisplayName("电话")]
        public string Tel { get; set; }
        [DisplayName("地址")]
        public string Address { get; set; }
    }
}
