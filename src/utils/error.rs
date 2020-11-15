use console::{style, Style, Term};
use lazy_static::lazy_static;
use thiserror::Error;

use std::io;

lazy_static! {
    pub static ref TERM_ERR: Term = Term::stderr();
    pub static ref TERM_OUT: Term = Term::stdout();
    static ref YELLOW: Style = Style::new().for_stderr().yellow();
    pub static ref GREEN: Style = Style::new().for_stderr().green();
    pub static ref MAGENTA: Style = Style::new().for_stderr().magenta();
}

macro_rules! _info {
    ($desc:literal, $val:expr) => {
        $crate::utils::TERM_ERR.write_line(&format!(
            "{} {} {}",
            $crate::utils::GREEN.apply_to("info"),
            $crate::utils::MAGENTA.apply_to($desc),
            $val
        ))
    };
}

pub(crate) use _info as info;

pub type Result<T = ()> = std::result::Result<T, Error>;

#[derive(Error, Debug)]
pub enum Error {
    #[error("{0}")]
    Serde(#[from] serde_json::Error),
    #[error("{0}")]
    Io(#[from] io::Error),
    #[error("cannot convert command output to string, {0}")]
    FromUtf8(#[from] std::string::FromUtf8Error),
}

impl Error {
    pub fn print_err(self) -> io::Result<()> {
        self.print(&TERM_ERR)
    }

    fn color(self) -> Self {
        match self {
            _ => self,
        }
    }

    pub fn print(self, term: &Term) -> io::Result<()> {
        term.write_str(&format!("{}: ", style("error").for_stderr().red().bold()))?;

        let msg = format!("{}", self.color());

        term.write_line(&msg)?;
        term.flush()
    }
}
