// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     TSTypesJenny
//     PluginTSTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

import * as common from '@grafana/schema';

export interface Options extends common.SingleStatBaseOptions {
  displayMode: common.BarGaugeDisplayMode;
  minVizHeight: number;
  minVizWidth: number;
  showUnfilled: boolean;
  valueMode: common.BarGaugeValueMode;
  valuePlacement: common.BarGaugeValuePlacement;
}

export const defaultOptions: Partial<Options> = {
  displayMode: common.BarGaugeDisplayMode.Gradient,
  minVizHeight: 10,
  minVizWidth: 0,
  showUnfilled: true,
  valueMode: common.BarGaugeValueMode.Color,
  valuePlacement: common.BarGaugeValuePlacement.Auto,
};
