using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class AttrValue
    {
        public long AttrValueID { get; set; }
        [DisplayName("名称")]
        public string AttrValueName { get; set; }
    }
}
