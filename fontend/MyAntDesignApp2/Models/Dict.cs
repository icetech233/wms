using System.ComponentModel;

namespace MyAntDesignApp2.Models
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
}
