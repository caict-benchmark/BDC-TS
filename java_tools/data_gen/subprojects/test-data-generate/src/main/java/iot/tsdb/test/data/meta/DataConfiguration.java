package iot.tsdb.test.data.meta;

public class DataConfiguration {
    public static String CSV_SPLITOR = ",";
    public static String DI_YA = "diya";

    public static String PROVINCE = "province";
    public static String DISTRICT = "district";
    public static String SYSTEM = "system";
    public static String MPID = "mpid";
    public static String CUSERID = "cuserid";
    public static String BJLX = "bjlx";
    public static String LINE = "line";
    public static String AREA = "area";

    public static String[] province = {"gd", "gx", "hn", "gz", "yn", "gz", "sz"};
    public static String[][] districts = {
            {"zhuhai", "shantou", "foshan", "shaoguan", "zhanjiang", "zhaoqing", "jiangmen", "maoming", "huizhou",
                    "meizhou", "shanwei", "heyuan", "yangjiang", "qingyuan", "dongguan", "zhongshan", "chouzhou",
                    "jieyang", "yunfu"},
            {"nanning", "liuzhou", "guilin", "wuzhou", "beihai", "fangchenggang", "qinzhou", "guigang", "yulin",
                    "baise", "hezhou", "hechi", "laibin", "chongzuo"},
            {"haikou", "sanya", "sansha", "danzhou", "wuzhishan", "wenchang", "qionghai", "wangning", "dongfang",
                    "anding", "tunchang", "chengmai", "lingao", "baisha", "changjiang", "ledong", "lingshui",
                    "baoting", "qiongzhong", "yangpu"},
            {"guiyang", "zunyi", "liupanshui", "anshun", "tongren", "bijie", "qianxinan", "qiandongnan", "qiannan"},
            {"kunming", "qujing", "yuxi", "shaotong", "baoshan", "lijiang", "puer", "lincang", "dehong", "nujiang",
                    "diqing", "dali", "chuxiong", "honghe", "wenshan", "xishuangbanna"},
            {"yuexiu", "haizhu", "liwan", "tianhe", "baiyun", "huangpu", "nansha", "fanyu", "huadu", "zengcheng",
                    "conghua"},
            {"futian", "luohu", "yantian", "nanshan", "baoan", "longgang", "longhua", "pingshan", "guangming", "dapeng"}
    };

    public static String[] fields = {
            "pos_pe_total", "pos_pe_peak", "pos_pe_flat", "pos_pe_valley", "pos_pe_tine",
            "pos_qe_total", "pos_qe_peak", "pos_qe_flat", "pos_qe_valley", "pos_qe_tine",
            "rev_pe_total", "rev_pe_peak", "rev_pe_flat", "rev_pe_valley", "rev_pe_tine",
            "rev_qe_total", "rev_qe_peak", "rev_qe_flat", "rev_qe_valley", "rev_qe_tine"
    };

    public static String getProvince(int userId) {
        return province[userId % province.length];
    }

    public static String getDistrict(int userId) {
        int index  = userId % province.length;
        return districts[index][userId % districts[index].length];
    }

    public static String getMpid(int userId) {
        return "00000";
    }

    public static String getSystem(int userId) {
        return "TMR";
    }

    public static String getBjlx(int userId) {
        return "3";
    }

    public static String getLine(int userId) {
        return "line_" + userId % 4000;
    }

    public static String getArea(int userId) {
        return "area_" + userId % 45000;
    }
}
