package resources

import (
	"github.com/rs/zerolog/log"
	"github.com/steadybit/action-kit/go/action_kit_api/v2"
)

const (
	actionIDs        = "com.github.steadybit.extension_host.host"
	stressCPUIcon    = "data:image/svg+xml,%3Csvg%20width%3D%2224%22%20height%3D%2224%22%20viewBox%3D%220%200%2024%2024%22%20fill%3D%22none%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M5.25%204.5C4.83579%204.5%204.5%204.83579%204.5%205.25V18.75C4.5%2019.1642%204.83579%2019.5%205.25%2019.5H18.75C19.1642%2019.5%2019.5%2019.1642%2019.5%2018.75V5.25C19.5%204.83579%2019.1642%204.5%2018.75%204.5H5.25ZM3%205.25C3%204.00736%204.00736%203%205.25%203H18.75C19.9926%203%2021%204.00736%2021%205.25V18.75C21%2019.9926%2019.9926%2021%2018.75%2021H5.25C4.00736%2021%203%2019.9926%203%2018.75V5.25Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M12%200.75C12.4142%200.75%2012.75%201.08579%2012.75%201.5V3.75C12.75%204.16421%2012.4142%204.5%2012%204.5C11.5858%204.5%2011.25%204.16421%2011.25%203.75V1.5C11.25%201.08579%2011.5858%200.75%2012%200.75Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M6.75%200.75C7.16421%200.75%207.5%201.08579%207.5%201.5V3.75C7.5%204.16421%207.16421%204.5%206.75%204.5C6.33579%204.5%206%204.16421%206%203.75V1.5C6%201.08579%206.33579%200.75%206.75%200.75Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M17.25%200.75C17.6642%200.75%2018%201.08579%2018%201.5V3.75C18%204.16421%2017.6642%204.5%2017.25%204.5C16.8358%204.5%2016.5%204.16421%2016.5%203.75V1.5C16.5%201.08579%2016.8358%200.75%2017.25%200.75Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M12%2019.5C12.4142%2019.5%2012.75%2019.8358%2012.75%2020.25V22.5C12.75%2022.9142%2012.4142%2023.25%2012%2023.25C11.5858%2023.25%2011.25%2022.9142%2011.25%2022.5V20.25C11.25%2019.8358%2011.5858%2019.5%2012%2019.5Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M6.75%2019.5C7.16421%2019.5%207.5%2019.8358%207.5%2020.25V22.5C7.5%2022.9142%207.16421%2023.25%206.75%2023.25C6.33579%2023.25%206%2022.9142%206%2022.5V20.25C6%2019.8358%206.33579%2019.5%206.75%2019.5Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M17.25%2019.5C17.6642%2019.5%2018%2019.8358%2018%2020.25V22.5C18%2022.9142%2017.6642%2023.25%2017.25%2023.25C16.8358%2023.25%2016.5%2022.9142%2016.5%2022.5V20.25C16.5%2019.8358%2016.8358%2019.5%2017.25%2019.5Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M19.5%2012C19.5%2011.5858%2019.8358%2011.25%2020.25%2011.25H22.5C22.9142%2011.25%2023.25%2011.5858%2023.25%2012C23.25%2012.4142%2022.9142%2012.75%2022.5%2012.75H20.25C19.8358%2012.75%2019.5%2012.4142%2019.5%2012Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M19.5%2017.25C19.5%2016.8358%2019.8358%2016.5%2020.25%2016.5H22.5C22.9142%2016.5%2023.25%2016.8358%2023.25%2017.25C23.25%2017.6642%2022.9142%2018%2022.5%2018H20.25C19.8358%2018%2019.5%2017.6642%2019.5%2017.25Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M19.5%206.75C19.5%206.33579%2019.8358%206%2020.25%206H22.5C22.9142%206%2023.25%206.33579%2023.25%206.75C23.25%207.16421%2022.9142%207.5%2022.5%207.5H20.25C19.8358%207.5%2019.5%207.16421%2019.5%206.75Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M0.75%2012C0.75%2011.5858%201.08579%2011.25%201.5%2011.25H3.75C4.16421%2011.25%204.5%2011.5858%204.5%2012C4.5%2012.4142%204.16421%2012.75%203.75%2012.75H1.5C1.08579%2012.75%200.75%2012.4142%200.75%2012Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M0.75%2017.25C0.75%2016.8358%201.08579%2016.5%201.5%2016.5H3.75C4.16421%2016.5%204.5%2016.8358%204.5%2017.25C4.5%2017.6642%204.16421%2018%203.75%2018H1.5C1.08579%2018%200.75%2017.6642%200.75%2017.25Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M0.75%206.75C0.75%206.33579%201.08579%206%201.5%206H3.75C4.16421%206%204.5%206.33579%204.5%206.75C4.5%207.16421%204.16421%207.5%203.75%207.5H1.5C1.08579%207.5%200.75%207.16421%200.75%206.75Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M8.25%207.5C7.83579%207.5%207.5%207.83579%207.5%208.25V15.75C7.5%2016.1642%207.83579%2016.5%208.25%2016.5H15.75C16.1642%2016.5%2016.5%2016.1642%2016.5%2015.75V8.25C16.5%207.83579%2016.1642%207.5%2015.75%207.5H8.25ZM6%208.25C6%207.00736%207.00736%206%208.25%206H15.75C16.9926%206%2018%207.00736%2018%208.25V15.75C18%2016.9926%2016.9926%2018%2015.75%2018H8.25C7.00736%2018%206%2016.9926%206%2015.75V8.25Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M11.25%2014.25C11.25%2013.8358%2011.5858%2013.5%2012%2013.5H14.25C14.6642%2013.5%2015%2013.8358%2015%2014.25C15%2014.6642%2014.6642%2015%2014.25%2015H12C11.5858%2015%2011.25%2014.6642%2011.25%2014.25Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3C%2Fsvg%3E%0A"
	stressIOIcon     = "data:image/svg+xml,%3Csvg%20width%3D%2224%22%20height%3D%2224%22%20viewBox%3D%220%200%2024%2024%22%20fill%3D%22none%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%0A%3Cpath%20d%3D%22M18.375%2017.625C18.3008%2017.625%2018.2283%2017.647%2018.1667%2017.6882C18.105%2017.7294%2018.0569%2017.788%2018.0285%2017.8565C18.0002%2017.925%2017.9927%2018.0004%2018.0072%2018.0732C18.0217%2018.1459%2018.0574%2018.2127%2018.1098%2018.2652C18.1623%2018.3176%2018.2291%2018.3533%2018.3018%2018.3678C18.3746%2018.3823%2018.45%2018.3748%2018.5185%2018.3465C18.587%2018.3181%2018.6456%2018.27%2018.6868%2018.2083C18.728%2018.1467%2018.75%2018.0742%2018.75%2018C18.75%2017.9005%2018.7105%2017.8052%2018.6402%2017.7348C18.5698%2017.6645%2018.4745%2017.625%2018.375%2017.625Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20d%3D%22M15%2017.625C14.9258%2017.625%2014.8533%2017.647%2014.7917%2017.6882C14.73%2017.7294%2014.6819%2017.788%2014.6535%2017.8565C14.6252%2017.925%2014.6177%2018.0004%2014.6322%2018.0732C14.6467%2018.1459%2014.6824%2018.2127%2014.7348%2018.2652C14.7873%2018.3176%2014.8541%2018.3533%2014.9268%2018.3678C14.9996%2018.3823%2015.075%2018.3748%2015.1435%2018.3465C15.212%2018.3181%2015.2706%2018.27%2015.3118%2018.2083C15.353%2018.1467%2015.375%2018.0742%2015.375%2018C15.375%2017.9005%2015.3355%2017.8052%2015.2652%2017.7348C15.1948%2017.6645%2015.0995%2017.625%2015%2017.625Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M14.375%2017.0646C14.56%2016.941%2014.7775%2016.875%2015%2016.875C15.2984%2016.875%2015.5845%2016.9935%2015.7955%2017.2045C16.0065%2017.4155%2016.125%2017.7016%2016.125%2018C16.125%2018.2225%2016.059%2018.44%2015.9354%2018.625C15.8118%2018.81%2015.6361%2018.9542%2015.4305%2019.0394C15.225%2019.1245%2014.9988%2019.1468%2014.7805%2019.1034C14.5623%2019.06%2014.3618%2018.9528%2014.2045%2018.7955C14.0472%2018.6382%2013.94%2018.4377%2013.8966%2018.2195C13.8532%2018.0012%2013.8755%2017.775%2013.9606%2017.5695C14.0458%2017.3639%2014.19%2017.1882%2014.375%2017.0646ZM15.1435%2018.3465C15.1661%2018.3371%2015.1878%2018.3255%2015.2083%2018.3118C15.2495%2018.2843%2015.2846%2018.2491%2015.3118%2018.2083C15.3254%2018.188%2015.337%2018.1663%2015.3465%2018.1435C15.3654%2018.0978%2015.375%2018.049%2015.375%2018C15.375%2017.9756%2015.3726%2017.951%2015.3678%2017.9268C15.3533%2017.8541%2015.3176%2017.7873%2015.2652%2017.7348C15.2127%2017.6824%2015.1459%2017.6467%2015.0732%2017.6322C15.0489%2017.6274%2015.0244%2017.625%2015%2017.625C14.951%2017.625%2014.9022%2017.6346%2014.8565%2017.6535C14.8337%2017.663%2014.812%2017.6746%2014.7917%2017.6882C14.7509%2017.7154%2014.7157%2017.7505%2014.6882%2017.7917C14.6745%2017.8122%2014.6629%2017.8339%2014.6535%2017.8565C14.6348%2017.9018%2014.625%2017.9505%2014.625%2018C14.625%2018.0247%2014.6274%2018.0492%2014.6322%2018.0732C14.6467%2018.1459%2014.6824%2018.2127%2014.7348%2018.2652C14.7873%2018.3176%2014.8541%2018.3533%2014.9268%2018.3678C14.9508%2018.3726%2014.9753%2018.375%2015%2018.375C15.0495%2018.375%2015.0982%2018.3652%2015.1435%2018.3465Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M5.25%2014.25C4.25544%2014.25%203.30161%2014.6451%202.59835%2015.3484C1.89509%2016.0516%201.5%2017.0054%201.5%2018C1.5%2018.9946%201.89509%2019.9484%202.59835%2020.6516C3.30161%2021.3549%204.25544%2021.75%205.25%2021.75H18.75C19.7446%2021.75%2020.6984%2021.3549%2021.4016%2020.6516C22.1049%2019.9484%2022.5%2018.9946%2022.5%2018C22.5%2017.0054%2022.1049%2016.0516%2021.4016%2015.3484C20.6984%2014.6451%2019.7446%2014.25%2018.75%2014.25H5.25ZM1.53769%2014.2877C2.52226%2013.3031%203.85761%2012.75%205.25%2012.75H18.75C20.1424%2012.75%2021.4777%2013.3031%2022.4623%2014.2877C23.4469%2015.2723%2024%2016.6076%2024%2018C24%2019.3924%2023.4469%2020.7277%2022.4623%2021.7123C21.4777%2022.6969%2020.1424%2023.25%2018.75%2023.25H5.25C3.85761%2023.25%202.52226%2022.6969%201.53769%2021.7123C0.553123%2020.7277%200%2019.3924%200%2018C0%2016.6076%200.553123%2015.2723%201.53769%2014.2877Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M6.87806%200.75C6.87804%200.75%206.87808%200.75%206.87806%200.75H17.123C17.9685%200.750211%2018.7894%201.03617%2019.4519%201.56146C20.1145%202.08673%2020.5801%202.82048%2020.7732%203.64364C20.7732%203.6436%2020.7732%203.64368%2020.7732%203.64364L23.8612%2016.8016C23.9558%2017.2049%2023.7056%2017.6085%2023.3024%2017.7032C22.8991%2017.7978%2022.4955%2017.5476%2022.4008%2017.1444L19.3128%203.98636C19.197%203.49244%2018.9176%203.05205%2018.5201%202.73688C18.1226%202.42174%2017.6303%202.25017%2017.123%202.25C17.1229%202.25%2017.1231%202.25%2017.123%202.25H6.878C6.37055%202.24996%205.87792%202.42145%205.48022%202.73664C5.08253%203.05183%204.80306%203.4922%204.68719%203.98625L1.59916%2017.1444C1.50452%2017.5476%201.1009%2017.7978%200.697641%2017.7032C0.294384%2017.6085%200.0441994%2017.2049%200.138838%2016.8016L3.22681%203.64375C3.2268%203.64379%203.22682%203.64371%203.22681%203.64375C3.41994%202.82038%203.88574%202.08637%204.54854%201.56107C5.21135%201.03577%206.03233%200.749943%206.87806%200.75Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M4.5%2018C4.5%2017.5858%204.83579%2017.25%205.25%2017.25H9C9.41421%2017.25%209.75%2017.5858%209.75%2018C9.75%2018.4142%209.41421%2018.75%209%2018.75H5.25C4.83579%2018.75%204.5%2018.4142%204.5%2018Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3C%2Fsvg%3E%0A"
	stressMemoryIcon = "data:image/svg+xml,%3Csvg%20width%3D%2224%22%20height%3D%2224%22%20viewBox%3D%220%200%2024%2024%22%20fill%3D%22none%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M17.1063%201.49823C16.9943%201.49823%2016.8834%201.52037%2016.7799%201.56338C16.6765%201.6064%2016.5826%201.66943%2016.5036%201.74886L16.5019%201.75054L10.5432%207.70453L10.5609%207.78931C10.6379%208.16975%2010.6196%208.56331%2010.5077%208.93498C10.3958%209.30665%2010.1938%209.64491%209.91967%209.91966C9.6455%2010.1944%209.30767%2010.3971%208.93624%2010.5098C8.56481%2010.6225%208.17129%2010.6416%207.79069%2010.5654L7.78392%2010.5641L7.70419%2010.5473L1.75019%2016.5023L1.74867%2016.5038C1.66924%2016.5828%201.60621%2016.6767%201.5632%2016.7801C1.52019%2016.8836%201.49805%2016.9945%201.49805%2017.1065C1.49805%2017.2185%201.52019%2017.3294%201.5632%2017.4329C1.60621%2017.5363%201.66924%2017.6302%201.74867%2017.7092L1.75015%2017.7107L6.29061%2022.2511C6.3696%2022.3306%206.46351%2022.3936%206.56695%2022.4366C6.67038%2022.4796%206.78129%2022.5018%206.89331%2022.5018C7.00534%2022.5018%207.11625%2022.4796%207.21968%2022.4366C7.32312%2022.3936%207.41703%2022.3306%207.49602%2022.2511L7.49748%2022.2497L22.2495%207.49767L22.251%207.4962C22.3304%207.41721%2022.3934%207.3233%2022.4364%207.21987C22.4794%207.11644%2022.5016%207.00552%2022.5016%206.8935C22.5016%206.78147%2022.4794%206.67056%2022.4364%206.56713C22.3934%206.4637%2022.3304%206.36978%2022.251%206.29079L22.2495%206.28933L17.7105%201.75033L17.709%201.74886C17.63%201.66943%2017.5361%201.60639%2017.4327%201.56338C17.3292%201.52037%2017.2183%201.49823%2017.1063%201.49823ZM16.204%200.178361C16.49%200.0594469%2016.7966%20-0.00177002%2017.1063%20-0.00177002C17.416%20-0.00177002%2017.7227%200.0594468%2018.0086%200.178361C18.2942%200.297124%2018.5536%200.4711%2018.7718%200.690304L18.7726%200.691138L23.3087%205.2272L23.3094%205.22795C23.5287%205.44618%2023.7027%205.70555%2023.8214%205.99119C23.9404%206.27715%2024.0016%206.5838%2024.0016%206.8935C24.0016%207.2032%2023.9404%207.50984%2023.8214%207.79581C23.7027%208.08145%2023.5287%208.34082%2023.3094%208.55905L23.3087%208.55979L8.55961%2023.3089L8.55886%2023.3096C8.34063%2023.5289%208.08126%2023.7029%207.79563%2023.8216C7.50966%2023.9405%207.20301%2024.0018%206.89331%2024.0018C6.58362%2024.0018%206.27697%2023.9405%205.991%2023.8216C5.70537%2023.7029%205.446%2023.5289%205.22777%2023.3096L5.22702%2023.3089L0.690955%2018.7728L0.690121%2018.772C0.470917%2018.5538%200.296941%2018.2944%200.178177%2018.0088C0.0592636%2017.7228%20-0.00195312%2017.4162%20-0.00195312%2017.1065C-0.00195312%2016.7968%200.0592638%2016.4901%200.178177%2016.2042C0.296919%2015.9186%200.470851%2015.6593%200.689999%2015.4412L0.690955%2015.4402L6.93044%209.19971C7.1095%209.02062%207.36684%208.94399%207.6147%208.99595L8.08778%209.09513C8.22508%209.12212%208.36692%209.115%208.50084%209.07437C8.6357%209.03347%208.75835%208.95987%208.85789%208.86012C8.95742%208.76037%209.03076%208.63756%209.07138%208.50263C9.11179%208.36839%209.11857%208.22629%209.09113%208.08884L8.99177%207.61488C8.93979%207.36694%209.01649%207.10952%209.1957%206.93045L15.44%200.691138L15.4411%200.690074C15.6592%200.470977%2015.9185%200.297082%2016.204%200.178361Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M7.49725%2015.4419C7.79001%2015.1489%208.26489%2015.1487%208.55791%2015.4414L9.69291%2016.5754C9.83364%2016.716%209.91275%2016.9068%209.91281%2017.1058C9.91288%2017.3047%209.8339%2017.4955%209.69326%2017.6362L7.42426%2019.9062C7.28362%2020.0469%207.09284%2020.126%206.8939%2020.126C6.69496%2020.126%206.50416%2020.047%206.36348%2019.9063L5.22848%2018.7713C4.93559%2018.4784%204.93559%2018.0036%205.22848%2017.7107C5.52138%2017.4178%205.99625%2017.4178%206.28914%2017.7107L6.8937%2018.3152L8.10204%2017.1063L7.49772%2016.5026C7.2047%2016.2098%207.20449%2015.7349%207.49725%2015.4419Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M17.7105%205.22867C18.0034%204.93577%2018.4782%204.93577%2018.7711%205.22867L19.9061%206.36367C20.0468%206.50434%2020.1258%206.69514%2020.1258%206.89408C20.1258%207.09302%2020.0467%207.2838%2019.906%207.42444L17.636%209.69344C17.4953%209.83409%2017.3045%209.91306%2017.1056%209.913C16.9066%209.91293%2016.7159%209.83383%2016.5753%209.69309L15.4413%208.55809C15.1485%208.26507%2015.1487%207.7902%2015.4417%207.49743C15.7347%207.20467%2016.2096%207.20488%2016.5024%207.4979L17.1062%208.10222L18.315%206.89388L17.7105%206.28933C17.4176%205.99643%2017.4176%205.52156%2017.7105%205.22867Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M13.1715%209.76767C13.4644%209.47477%2013.9392%209.47477%2014.2321%209.76767L15.3671%2010.9027C15.5078%2011.0433%2015.5868%2011.2341%2015.5868%2011.433C15.5868%2011.6319%2015.5078%2011.8227%2015.3671%2011.9633L11.9631%2015.3673C11.8225%2015.508%2011.6317%2015.587%2011.4328%2015.587C11.2339%2015.587%2011.0431%2015.508%2010.9025%2015.3673L9.76748%2014.2323C9.47459%2013.9394%209.47459%2013.4646%209.76748%2013.1717C10.0604%2012.8788%2010.5353%2012.8788%2010.8281%2013.1717L11.4328%2013.7763L13.7762%2011.433L13.1715%2010.8283C12.8786%2010.5354%2012.8786%2010.0606%2013.1715%209.76767Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M1.82472%2014.3064C2.11774%2014.0137%202.59261%2014.0139%202.88538%2014.3069L4.01938%2015.4419C4.31214%2015.7349%204.31193%2016.2098%204.01891%2016.5026C3.72589%2016.7953%203.25101%2016.7951%202.95825%2016.5021L1.82425%2015.3671C1.53149%2015.0741%201.5317%2014.5992%201.82472%2014.3064Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M4.09348%2012.0367C4.38638%2011.7438%204.86125%2011.7438%205.15414%2012.0367L6.28914%2013.1717C6.58204%2013.4646%206.58204%2013.9394%206.28914%2014.2323C5.99625%2014.5252%205.52138%2014.5252%205.22848%2014.2323L4.09348%2013.0973C3.80059%2012.8044%203.80059%2012.3296%204.09348%2012.0367Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M12.0365%204.09367C12.3294%203.80077%2012.8043%203.80077%2013.0971%204.09367L14.2321%205.22867C14.525%205.52156%2014.525%205.99643%2014.2321%206.28933C13.9392%206.58222%2013.4644%206.58222%2013.1715%206.28933L12.0365%205.15433C11.7436%204.86143%2011.7436%204.38656%2012.0365%204.09367Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3Cpath%20fill-rule%3D%22evenodd%22%20clip-rule%3D%22evenodd%22%20d%3D%22M14.3063%201.8249C14.599%201.53188%2015.0739%201.53167%2015.3669%201.82443L16.5019%202.95843C16.7949%203.2512%2016.7951%203.72607%2016.5024%204.01909C16.2096%204.31212%2015.7347%204.31232%2015.4417%204.01956L14.3067%202.88556C14.0137%202.5928%2014.0135%202.11792%2014.3063%201.8249Z%22%20fill%3D%22%231D2632%22%2F%3E%0A%3C%2Fsvg%3E%0A"
)

type StressActionState struct {
	StressNGArgs []string
	Pid          int
}

func start(state *StressActionState) (*action_kit_api.StartResult, error) {
	pid, err := StartStressNG(state.StressNGArgs)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start stress-ng")
		return nil, err
	}
	log.Info().Int("Pid", pid).Msg("Started stress-ng")
	state.Pid = pid
	return nil, nil
}

func stop(state *StressActionState) (*action_kit_api.StopResult, error) {
	if state.Pid != 0 {
		log.Info().Int("Pid", state.Pid).Msg("Stopping stress-ng")
		err := StopStressNG(state.Pid)
		if err != nil {
			log.Error().Err(err).Int("Pid", state.Pid).Msg("Failed to stop stress-ng")
			return nil, err
		}
		state.Pid = 0
	}
	return nil, nil
}
