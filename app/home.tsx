import { ScreenContent } from "components/screen/ScreenContent";
import StudyPerformanceChart from "components/ui/StudyPerformanceChart"; 
import { Text, View } from "react-native";
import MaterialCommunityIcons from '@expo/vector-icons/MaterialCommunityIcons';
import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import FontAwesome5 from '@expo/vector-icons/FontAwesome5';
import Entypo from '@expo/vector-icons/Entypo';

export default function Home() {
    // Placholder variables
    const hoursStudied = 2;
    const minutesStudied = 45;
    const tasksTotal = 12;
    const tasksCompleted = 9;    
    const mockStudyLog = {
        '2026-08-05': 40,
        '2026-08-06': 25,
        '2026-08-07': 60,
        '2026-08-08': 0,
        '2026-08-09': 55,
        '2026-08-10': 30,
        '2026-08-11': 45,
    };
    
    return (
        <ScreenContent title="Home" path="/home">
            <Text className="text-2xl font-bold text-gray-800 dark:text-white">Meu Dashboard</Text>
            {/* Cards de informações */}
            <View className="flex-row justify-between">
                <View className="mt-4 mr-2 flex-row items-center justify-between rounded-lg bg-white p-4 dark:bg-gray-800">
                    {/* Textos */}
                    <View>
                        <Text className="text-sm text-gray-500 dark:text-gray-400">
                            Tempo Estudado
                        </Text>

                        <Text className="text-3xl font-bold text-gray-800 dark:text-white">
                            {hoursStudied}h {minutesStudied}m
                        </Text>
                    </View>

                    {/* Ícone */}
                    <View className="h-18 w-18 px-2 py-4 ml-5 items-center justify-center rounded-full bg-primary-200 dark:bg-primary-900">
                        <MaterialCommunityIcons name="clock" size={28} color="#0284c7" />
                    </View>
                </View>
                <View className="mt-4 flex-row ml-2 items-center justify-between rounded-lg bg-white p-4 dark:bg-gray-800">
                    {/* Textos */}
                    <View className="mr-4">
                        <Text className="text-sm text-gray-500 dark:text-gray-400">
                            Tarefas
                        </Text>

                        <Text className="text-3xl font-bold text-gray-800 dark:text-white">
                            {tasksCompleted} / {tasksTotal}
                        </Text>
                    </View>

                    {/* Ícone */}
                    <View className="h-18 w-18 px-2 py-4 ml-5 items-center justify-center rounded-full bg-green-200 dark:bg-green-900">
                        <MaterialCommunityIcons name="check-circle" size={28} color="#16A34A" />
                    </View>
                </View>
            </View>
            {/* Gráfico */}
            <StudyPerformanceChart studyLog={mockStudyLog} />
            {/* Seção 1 */}
            <View className="mt-4 flex-row justify-between">
                <View className="bg-white pt-5 pb-5 pr-7 pl-7 rounded-lg dark:bg-gray-800">
                    <View className="items-center rounded-full py-8 px-[5px] bg-purple-100 dark:bg-purple-900 " >
                        {/* <MaterialIcons name="calendar-month" size={24} color="#9333EA" /> */}
                        {/* <FontAwesome5 name="calendar-alt" size={24} color="#9333EA" /> */}
                        <Entypo name="calendar" size={24} color="#9333EA" />
                    </View>
                    <Text className="mt-2 text-md font-medium text-gray-800 dark:text-white">
                        Calendário
                    </Text>
                </View>
                <View className="bg-white pt-5 pb-5 pr-8 pl-8 rounded-lg dark:bg-gray-800">
                    <View className="items-center rounded-full py-7 bg-blue-100 dark:bg-blue-900" >
                        <FontAwesome5 name="book" size={24} color="#2563EB" />
                    </View>
                    <Text className="mt-2 text-md font-medium text-gray-800 dark:text-white">
                        Resumos
                    </Text>
                </View>
                <View className="bg-white pt-5 pb-5 pr-10 pl-10 rounded-lg dark:bg-gray-800">
                    <View className="items-center rounded-full bg-green-200 dark:bg-green-900  py-7" >
                        <FontAwesome5 name="tasks" size={24} color="#16A34A" />                    
                    </View>
                    <Text className="mt-2 text-md font-medium text-gray-800 dark:text-white">
                        Tarefas
                    </Text>
                </View>
            </View>
            {/* Seção 2 */}
            <View className="mt-4 flex-row justify-between">
                <View className="bg-white pt-5 pb-5 pr-9 pl-9 rounded-lg dark:bg-gray-800">
                    <View className="items-center rounded-full bg-yellow-200 dark:bg-yellow-900  py-7" >
                        <Entypo name="archive" size={24} color="#CA8A04" />                    
                    </View>
                    <Text className="mt-2 text-md font-medium text-gray-800 dark:text-white">
                        Arquivos
                    </Text>
                </View>
            </View>
        </ScreenContent>
    );
}