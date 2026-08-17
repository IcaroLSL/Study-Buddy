import React, { useMemo, useState } from 'react';
import {
  View,
  Text,
  ScrollView,
  TouchableOpacity,
  Modal,
} from 'react-native';

/**
 * StudyPerformanceChart
 * Versão TypeScript + NativeWind (Tailwind) do gráfico "Seu Desempenho".
 * Usa as mesmas classes/cores do seu tailwind.config.js (primary-500 etc).
 */

export type RangeOption = 'Esta Semana' | 'Este Mês' | 'Este Ano';

export interface StudyLog {
  /** chave no formato 'YYYY-MM-DD' -> minutos estudados */
  [date: string]: number;
}

interface ChartItem {
  label: string;
  value: number;
}

export interface StudyPerformanceChartProps {
  studyLog?: StudyLog;
  initialRange?: RangeOption;
}

const RANGE_OPTIONS: RangeOption[] = ['Esta Semana', 'Este Mês', 'Este Ano'];
const WEEK_LABELS = ['Dom', 'Seg', 'Ter', 'Qua', 'Qui', 'Sex', 'Sáb'];
const MONTH_LABELS = [
  'Jan', 'Fev', 'Mar', 'Abr', 'Mai', 'Jun',
  'Jul', 'Ago', 'Set', 'Out', 'Nov', 'Dez',
];

const CHART_HEIGHT = 160; // h-40
const MAX_BAR_HEIGHT = 64;
const BAR_WIDTH = 40; // w-10
const BAR_GAP = 8; // gap-2

function pad2(n: number): string {
  return n.toString().padStart(2, '0');
}

function getItemsForRange(studyLog: StudyLog, range: RangeOption): ChartItem[] {
  const today = new Date();
  const items: ChartItem[] = [];

  if (range === 'Esta Semana') {
    const start = new Date(today);
    start.setDate(today.getDate() - today.getDay());
    for (let i = 0; i < 7; i++) {
      const current = new Date(start);
      current.setDate(start.getDate() + i);
      const key = current.toISOString().split('T')[0];
      const minutes = studyLog[key] || 0;
      items.push({ label: WEEK_LABELS[i], value: minutes });
    }
  }

  if (range === 'Este Mês') {
    const year = today.getFullYear();
    const month = today.getMonth();
    const daysInMonth = new Date(year, month + 1, 0).getDate();
    for (let i = 1; i <= daysInMonth; i++) {
      const date = `${year}-${pad2(month + 1)}-${pad2(i)}`;
      const minutes = studyLog[date] || 0;
      items.push({ label: i.toString(), value: minutes });
    }
  }

  if (range === 'Este Ano') {
    const year = today.getFullYear();
    for (let i = 0; i < 12; i++) {
      let total = 0;
      for (let d = 1; d <= 31; d++) {
        const date = `${year}-${pad2(i + 1)}-${pad2(d)}`;
        total += studyLog[date] || 0;
      }
      items.push({ label: MONTH_LABELS[i], value: total });
    }
  }

  return items;
}

export default function StudyPerformanceChart({
  studyLog = {},
  initialRange = 'Esta Semana',
}: StudyPerformanceChartProps) {
  const [range, setRange] = useState<RangeOption>(initialRange);
  const [pickerOpen, setPickerOpen] = useState(false);

  const items = useMemo(() => getItemsForRange(studyLog, range), [studyLog, range]);

  const isScrollable = range !== 'Esta Semana';
  const contentWidth = isScrollable ? items.length * (BAR_WIDTH + BAR_GAP) : undefined;

  return (
    <View className="bg-white dark:bg-gray-800 rounded-lg p-4  mt-4 shadow mb-6">
      <View className="flex-row justify-between items-center mb-3">
        <Text className="font-medium text-gray-900 dark:text-gray-50">Seu Desempenho</Text>

        <TouchableOpacity
          className="flex-row items-center gap-1 bg-gray-100 dark:bg-gray-700 rounded px-2 py-1"
          onPress={() => setPickerOpen(true)}
          activeOpacity={0.7}
        >
          <Text className="text-sm text-gray-900 dark:text-gray-50">{range}</Text>
          <Text className="text-xs text-gray-500 dark:text-gray-400">▾</Text>
        </TouchableOpacity>
      </View>

      <View className="bg-gray-100 dark:bg-gray-700 rounded px-3 py-2">
        {isScrollable ? (
          <ScrollView horizontal showsHorizontalScrollIndicator={false}>
            <ChartBars items={items} width={contentWidth} justify="flex-start" />
          </ScrollView>
        ) : (
          <ChartBars items={items} justify="space-between" fullWidth />
        )}
      </View>

      <RangePickerModal
        visible={pickerOpen}
        current={range}
        onClose={() => setPickerOpen(false)}
        onSelect={(value) => {
          setRange(value);
          setPickerOpen(false);
        }}
      />
    </View>
  );
}

interface ChartBarsProps {
  items: ChartItem[];
  width?: number;
  justify: 'flex-start' | 'space-between';
  fullWidth?: boolean;
}

function ChartBars({ items, width, justify, fullWidth }: ChartBarsProps) {
  return (
    <View
      className="flex-row items-end gap-2"
      style={{
        height: CHART_HEIGHT,
        justifyContent: justify,
        width: fullWidth ? '100%' : width,
      }}
    >
      {items.map((item, index) => {
        const barHeight = Math.min(MAX_BAR_HEIGHT, item.value * 2);
        return (
          <View
            key={`${item.label}-${index}`}
            className="items-center justify-end"
            style={{ width: BAR_WIDTH }}
          >
            <View className="w-full justify-end">
              <View
                className="w-full bg-primary-500 rounded-t"
                style={{ height: barHeight }}
              />
            </View>
            <Text
              className="text-xs mt-1 text-center text-gray-500 dark:text-gray-400"
              numberOfLines={1}
            >
              {item.label}
            </Text>
          </View>
        );
      })}
    </View>
  );
}

interface RangePickerModalProps {
  visible: boolean;
  current: RangeOption;
  onClose: () => void;
  onSelect: (value: RangeOption) => void;
}

function RangePickerModal({ visible, current, onClose, onSelect }: RangePickerModalProps) {
  return (
    <Modal visible={visible} transparent animationType="fade" onRequestClose={onClose}>
      <TouchableOpacity
        className="flex-1 justify-center items-center bg-black/40"
        activeOpacity={1}
        onPress={onClose}
      >
        <View className="bg-white dark:bg-gray-800 rounded-xl py-2 w-56">
          {RANGE_OPTIONS.map((option) => (
            <TouchableOpacity
              key={option}
              className="py-3 px-4"
              onPress={() => onSelect(option)}
            >
              <Text
                className={
                  option === current
                    ? 'text-base font-semibold text-primary-500'
                    : 'text-base text-gray-900 dark:text-gray-50'
                }
              >
                {option}
              </Text>
            </TouchableOpacity>
          ))}
        </View>
      </TouchableOpacity>
    </Modal>
  );
}

/*
Exemplo de uso:

import StudyPerformanceChart, { StudyLog } from './StudyPerformanceChart';

const studyLog: StudyLog = {
  '2026-08-10': 45,
  '2026-08-11': 90,
};

<StudyPerformanceChart studyLog={studyLog} />
*/