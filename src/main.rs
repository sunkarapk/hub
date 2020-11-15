use clap::{AppSettings, Clap};

use std::process::exit;

mod labels;

mod utils;

use utils::error::{TERM_ERR, TERM_OUT};

#[derive(Debug, Clap)]
enum Subcommand {
    Labels(labels::Labels),
}

#[derive(Debug, Clap)]
#[clap(
    version,
    global_setting(AppSettings::VersionlessSubcommands),
    global_setting(AppSettings::ColoredHelp)
)]
struct Opt {
    #[clap(subcommand)]
    subcommand: Subcommand,
}

fn main() {
    let opt = Opt::parse();

    let err = match opt.subcommand {
        Subcommand::Labels(x) => x.run(),
    }
    .err();

    let code = if let Some(error) = err {
        error.print_err().unwrap();
        1
    } else {
        0
    };

    TERM_ERR.flush().unwrap();
    TERM_OUT.flush().unwrap();

    exit(code)
}
