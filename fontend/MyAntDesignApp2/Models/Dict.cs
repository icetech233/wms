using System.ComponentModel;

namespace WmsApp.Models
{
    public class Dict : CommonDate
    {
        [DisplayName("ID")]
        public long DictID { get; set; }
        [DisplayName("业务类型")]
        public int BizType { get; set; }
        [DisplayName("选项")]
        public string DictName { get; set; }
        [DisplayName("值")]
        public string DictValue { get; set; }
    }

    public class Tenant : CommonDate
    {
        [DisplayName("ID")]
        public long TenantID { get; set; }
        [DisplayName("是否供应商")]
        public bool IsSupplier { get; set; }
        [DisplayName("是否客户")]
        public bool IsCustomer { get; set; }
        [DisplayName("统一社会信用代码")]
        public string TenantCode { get; set; }
        [DisplayName("企业名称")]
        public string TenantName { get; set; }
        [DisplayName("联系人")]
        public string Contacts { get; set; }
        [DisplayName("座机号")]
        public string Telephone { get; set; }
        [DisplayName("email")]
        public string Email { get; set; }
        [DisplayName("地址")]
        public string Address { get; set; }
    }
}
