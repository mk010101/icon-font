import { generateFonts, FontAssetType, OtherAssetType } from "fantasticon";

generateFonts({
  name: "icon-font",
  inputDir: "./icon-font/svg",
  outputDir: "./icon-font/font",
  fontTypes: [FontAssetType.WOFF2],
  assetTypes: [OtherAssetType.CSS, OtherAssetType.HTML],
  formatOptions: {
    format: true,
  },
  templates: {
    css: "./icon-font/template.hbs",
  },
  pathOptions: {},
  codepoints: {},
  fontHeight: 600,
  round: undefined,
  descent: 100,
  normalize: true,
  selector: null,
  tag: "div",
  prefix: "i",
  fontsUrl: "",
}).then((results) => console.log(results));
