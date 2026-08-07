import { ScreenContent } from 'components/ScreenContent';
import { StatusBar } from 'expo-status-bar';
import { SafeAreaProvider, SafeAreaView } from 'react-native-safe-area-context';

import '../global.css';

export default function App() {
  return (
    <SafeAreaProvider>
      <SafeAreaView className="flex-1 bg-slate-100">
        <ScreenContent />
        <StatusBar style="auto" />
      </SafeAreaView>
    </SafeAreaProvider>
  );
}
