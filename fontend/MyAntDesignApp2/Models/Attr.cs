using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class Attr
    {
        public long AttrID { get; set; }
        public string AttrName { get; set; }
        [DisplayName("创建")]
        public string CreatedAt { get; set; }
        [DisplayName("修改")]
        public string UpdatedAt { get; set; }
    }
}
