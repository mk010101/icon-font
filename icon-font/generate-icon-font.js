import { generateFonts, FontAssetType, OtherAssetType } from "fantasticon";

generateFonts({
  name: "a-icons",
  inputDir: "./icon-font/svg",
  outputDir: "./icon-font/font",
  fontTypes: [FontAssetType.WOFF, FontAssetType.TTF],
  assetTypes: [OtherAssetType.CSS, OtherAssetType.HTML],
  formatOptions: {
    format: true,
  },
  templates: {
    css: "./icon-font/template.hbs",
  },
  pathOptions: {},
  codepoints: {},
  fontHeight: 1024,
  round: undefined,
  descent: 118,
  normalize: true,
  selector: null,
  tag: "div",
  prefix: "icon",
  fontsUrl: "",
}).then((results) => console.log(results));
