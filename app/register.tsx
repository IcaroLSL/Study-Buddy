import { FontAwesome5 } from '@expo/vector-icons';
import { useRouter } from 'expo-router';
import { useMemo, useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  KeyboardAvoidingView,
  Platform,
  Pressable,
  ScrollView,
  Text,
  TextInput,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

const API_BASE_URL = 'http://localhost:8080';

type RegisterResponse = {
  message?: string;
  error?: string;
};

export default function RegisterScreen() {
  const router = useRouter();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isDark, setIsDark] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const canSubmit = useMemo(
    () => name.trim().length > 0 && email.trim().length > 0 && password.length > 0 && !isSubmitting,
    [name, email, password, isSubmitting]
  );

  const handleRegister = async () => {
    if (!canSubmit) {
      Alert.alert('Campos obrigatórios', 'Informe nome, email e senha para criar sua conta.');
      return;
    }

    setIsSubmitting(true);

    try {
      const response = await fetch(`${API_BASE_URL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name: name.trim(),
          email: email.trim(),
          password,
        }),
      });

      const data = (await response.json()) as RegisterResponse;

      if (!response.ok) {
        Alert.alert('Erro ao registrar', data.error ?? 'Erro ao registrar');
        return;
      }

      Alert.alert('Cadastro realizado', data.message ?? 'Usuário cadastrado com sucesso!', [
        { text: 'Entrar', onPress: () => router.replace('/') },
      ]);
    } catch (error) {
      console.error('Erro no cadastro:', error);
      Alert.alert('Erro de conexão', 'Não foi possível conectar com o servidor.');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <SafeAreaView className={isDark ? 'flex-1 bg-gray-950' : 'flex-1 bg-gray-100'}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1">
        <ScrollView
          contentContainerClassName="flex-grow justify-center px-4 py-8"
          keyboardShouldPersistTaps="handled">
          <View
            className={
              isDark
                ? 'rounded-3xl bg-gray-800 p-8 shadow-2xl'
                : 'rounded-3xl bg-white p-8 shadow-2xl'
            }>
            <View className="absolute right-4 top-4">
              <Pressable
                accessibilityLabel="Alternar tema"
                className={isDark ? 'rounded-full bg-gray-700 p-3' : 'rounded-full bg-gray-100 p-3'}
                onPress={() => setIsDark((current: boolean) => !current)}>
                <FontAwesome5
                  color={isDark ? '#d1d5db' : '#374151'}
                  name={isDark ? 'adjust' : 'moon'}
                  size={18}
                />
              </Pressable>
            </View>

            <View className="mb-6 items-center">
              <FontAwesome5 color={isDark ? '#3b82f6' : '#2563eb'} name="user-plus" size={40} />
              <Text
                className={
                  isDark
                    ? 'mt-4 text-2xl font-bold text-white'
                    : 'mt-4 text-2xl font-bold text-gray-800'
                }>
                Criar Conta
              </Text>
              <Text className={isDark ? 'mt-1 text-gray-300' : 'mt-1 text-gray-600'}>
                Comece sua jornada de estudos!
              </Text>
            </View>

            <View className="gap-4">
              <View>
                <Text className={isDark ? labelStyles.dark : labelStyles.light}>Nome</Text>
                <TextInput
                  autoCapitalize="words"
                  autoComplete="name"
                  className={isDark ? inputStyles.dark : inputStyles.light}
                  editable={!isSubmitting}
                  onChangeText={setName}
                  placeholder="Seu nome"
                  placeholderTextColor={isDark ? '#9ca3af' : '#6b7280'}
                  value={name}
                />
              </View>

              <View>
                <Text className={isDark ? labelStyles.dark : labelStyles.light}>Email</Text>
                <TextInput
                  autoCapitalize="none"
                  autoComplete="email"
                  className={isDark ? inputStyles.dark : inputStyles.light}
                  editable={!isSubmitting}
                  keyboardType="email-address"
                  onChangeText={setEmail}
                  placeholder="seu@email.com"
                  placeholderTextColor={isDark ? '#9ca3af' : '#6b7280'}
                  value={email}
                />
              </View>

              <View>
                <Text className={isDark ? labelStyles.dark : labelStyles.light}>Senha</Text>
                <TextInput
                  autoComplete="password"
                  className={isDark ? inputStyles.dark : inputStyles.light}
                  editable={!isSubmitting}
                  onChangeText={setPassword}
                  placeholder="••••••••"
                  placeholderTextColor={isDark ? '#9ca3af' : '#6b7280'}
                  secureTextEntry
                  value={password}
                />
              </View>

              <Pressable
                className={canSubmit ? buttonStyles.enabled : buttonStyles.disabled}
                disabled={!canSubmit}
                onPress={handleRegister}>
                {isSubmitting ? (
                  <ActivityIndicator color="#ffffff" />
                ) : (
                  <Text className="text-center font-medium text-white">Cadastrar</Text>
                )}
              </Pressable>
            </View>

            <View className="mt-6 flex-row justify-center">
              <Text className={isDark ? 'text-sm text-gray-400' : 'text-sm text-gray-600'}>
                Já tem uma conta?{' '}
              </Text>
              <Pressable onPress={() => router.replace('/')}>
                <Text className="text-sm font-medium text-blue-600">Entrar</Text>
              </Pressable>
            </View>
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

const labelStyles = {
  light: 'mb-1 text-sm font-medium text-gray-700',
  dark: 'mb-1 text-sm font-medium text-gray-300',
};

const inputStyles = {
  light:
    'w-full rounded-lg border border-gray-300 bg-white px-4 py-3 text-gray-900 focus:border-blue-500',
  dark: 'w-full rounded-lg border border-gray-600 bg-gray-700 px-4 py-3 text-white focus:border-blue-500',
};

const buttonStyles = {
  enabled: 'w-full rounded-lg bg-blue-600 px-4 py-3',
  disabled: 'w-full rounded-lg bg-blue-300 px-4 py-3',
};
