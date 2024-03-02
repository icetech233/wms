using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class Attr : CommonDate
    {
        public long AttrID { get; set; }
        public string AttrName { get; set; }
        public int ShowType { get; set; }
        [DisplayName("创建")]
        public string CreatedAt { get; set; }
        [DisplayName("修改")] 
        public string UpdatedAt { get; set; }
    }

    public class Menu : CommonDate
    {
        public long MenuID { get; set; }
        public long parentID { get; set; }
        public string Key { get; set; }
        public string Name { get; set; }
        public string Icon { get; set; }
        public string Path { get; set; }
    }


    public class CommonDate {
        [DisplayName("创建")]
        public string CreatedAt { get; set; }
        [DisplayName("修改")]
        public string UpdatedAt { get; set; }
    }
}
