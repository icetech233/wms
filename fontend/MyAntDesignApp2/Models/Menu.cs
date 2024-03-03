namespace MyAntDesignApp2.Models
{
    public class Menu : CommonDate
    {
        public long MenuID { get; set; }
        public long parentID { get; set; }
        public string Key { get; set; }
        public string Name { get; set; }
        public string Icon { get; set; }
        public string Path { get; set; }
    }
}
