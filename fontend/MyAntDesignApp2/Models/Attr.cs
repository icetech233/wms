using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class Attr : CommonDate
    {
        public long AttrID { get; set; }
        public string AttrName { get; set; }
        public int ShowType { get; set; }
        [DisplayName("值")]
        public AttrValue[] Val { get; set; }
    }

    public class AttrValue : CommonDate
    {
        public long AttrValueID { get; set; }
        [DisplayName("名称")]
        public object AttrValueName { get; set; }
    }
}
