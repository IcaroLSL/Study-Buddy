import { View, Text, Pressable } from "react-native"

interface CardTechniqueProps {
    title: string,
    description: string,
    cor: string,
    paddingX: string,
    vertical?: boolean,
    paddingY?: string,
    icon: React.ReactNode,
    redirect: () => void
}

export default function CardTechnique({ title, description, cor, paddingX, paddingY = "1", icon, redirect }: CardTechniqueProps) {
    return(
        <View className="bg-white dark:bg-gray-800 rounded-lg p-4 shadow mt-4">
            <View className="flex-row">
                <View className={`px-${paddingX} py-${paddingY} rounded-full items-center justify-center bg-${cor}-200 dark:bg-${cor}-900 mr-3 mt-2 mb-2`}>
                    {icon}
                </View>
                <View className="flex-1 min-w-0">
                    <Text className="font-medium mt-1 text-2xl text-black dark:text-white" style={{ flexShrink: 1 }}>
                        {title}
                    </Text>
                    <Text className="text-lg text-gray-500">
                        {description}
                    </Text>
                    <Pressable className="mt-2">
                        <Text className="text-primary-600 text-lg font-semibold">
                            Iniciar
                        </Text>
                    </Pressable>
                </View>
            </View>
        </View>
    )
}