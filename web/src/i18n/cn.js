import chineseMessages from 'ra-language-chinese';

export default {
    ...chineseMessages,
    pos: {
        search: '搜索',
        configuration: 'Configuration',
        language: 'Language',
        theme: {
            name: 'Theme',
            light: 'Light',
            dark: 'Dark',
        },
        dashboard: {
            monthly_revenue: 'Monthly Revenue',
            new_orders: 'New Orders',
            pending_reviews: 'Pending Reviews',
            new_customers: 'New Customers',
            pending_orders: 'Pending Orders',
            order: {
                items:
                    'by %{customer_name}, one item |||| by %{customer_name}, %{nb_items} items',
            },
            welcome: {
                title: 'Welcome to react-admin demo',
                subtitle:
                    "This is the admin of an imaginary poster shop. Fell free to explore and modify the data - it's local to your computer, and will reset each time you reload.",
                aor_button: 'react-admin site',
                demo_button: 'Source for this demo',
            },
        },
    },
    resources: {
        servers: {
            name: '服务器',
            fields: {
                name: '名称',
                alias: '别名',
                guid: 'GUID',
                serial: '序列号',
                modelName: '生产名称',
                osIpAddr: '操作系统业务地址',
                username: 'BMC用户名',
                password: 'BMC密码',
                ipaddr1: '上模块BMC地址',
                ipaddr2: '下模块BMC地址',
                primaryIp: '主模块IP',
            },
            tabs: {
                detail: '详细',
                status: '状态',
                setting: '设置',
            },
            page: {
                delete: 'Delete Customer',
            },
        },
        users: {
            name: '用户',
            fields: {
                name: '名称',
                username: '用户名',
                password: '密码',
                role: '权限',
            },
            tabs: {
                detail: '详细',
                status: '状态',
                setting: '设置',
            },
            page: {
                delete: 'Delete Customer',
            },
        },
        rooms: {
            name: '机房',
            fields: {
                name: '名称',
            },
            tabs: {
                detail: '详细',
            },
            page: {
                delete: 'Delete Customer',
            },
        },
        modules: {
            name: '模块',
            fields: {
                powerStatus: '系统电源状态',
                systemFtLed: '系统FT LED',
                systemFaultLed: '系统Fault LED',
                cpuFaultLed: 'CPU Fault LED',
                ioFaultLed: 'I/O Fault LED',
                safeToPull: 'safeToPull LED',
                primaryLed: '主模块LED',
                modulePowerLed: '模块电源 LED',
                moduleId: '模块ID LED',
                lcds: 'LCD',
                poh: '通电累计时间/小时',
            },
            tabs: {
                common: '共同状态',
            },
        },
        memorys: {
            name: '内存',
            tabs: {
                detail: '内存详细',
            },
            fields: {
                errorCorrection: '错误纠正',
                phyCapacity: '总物理内存',
                phyAlive: '可用物理内存',
                phyUsed: '已用物理内存',
                virtCapacity: '总虚拟内存',
                virtUsed: '已用虚拟内存',
                pageCapacity: '总内存页',
                pageUsed: '已用内存页',
                index: 'index',
                type: '类型',
                size: '容量',
                status: '状态',
            },
        },
        cpus: {
            name: 'CPU',
            fields: {
                phyNum: '物理CPU数量/颗',
                logicNum: '逻辑CPU线程数',
                threshold: 'Threshold',
                load_60: 'CPU负载/分钟',
                status: '状态',
            },
        },
        drivers: {
            name: '硬盘',
            fields: {
                letter: '盘符',
                capacity: '总容量',
                unusedCapacity: '可用容量',
                status: '状态',
                driveType: '驱动类型',
                label: 'label',
                serial: '序列号',
                type: '类型',
            },
        },
    },
    user: {
        choices: {
            role: {
                no: '无',
                user: '用户',
                admin: '管理员',
            },
        },
    },
};
