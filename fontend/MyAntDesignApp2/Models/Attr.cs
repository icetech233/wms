using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;

namespace WmsApp.Models
{
    public class Attr : CommonDate
    {
        public long AttrID { get; set; }
        public string AttrName { get; set; }
        public int ShowType { get; set; }
        [DisplayName("值")]
        public List<AttrValue> Val { get; set; }
    }
}
