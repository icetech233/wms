using System.Collections.Generic;
using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class Attr : CommonDate
    {
        public long AttrID { get; set; }
        public string AttrName { get; set; }
        public int ShowType { get; set; }
        [DisplayName("值")]
        public List<AttrValue> Val { get; set; }
    }

    public class AttrValue
    {
        public long AttrValueID { get; set; }
        [DisplayName("名称")]
        public string AttrValueName { get; set; }
    }
}
